/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/2 22:00
  */

package main

import (
	"sync"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

var (
	confPath = "config.yml"			// config 配置路径
)

type GroupsG struct {
	GroupsG		Groups
	lock 		sync.Mutex
}

type Groups struct {
	Group 		[]Group		`json:"Group"`
}

type Group struct {
	GroupName 	string		`json:"GroupName"`
	Previlege	[]Previlege	`json:"Previlege"`
}

type Previlege struct {
	TagName		string		`json:"TagName"`
	TagPrev 	string		`json:"TagPrev"`
}

// 判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filepath string) bool {
	var exist = true
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 写入文件
func writeFile(content, filePath string) error {
	var f *os.File
	var err error
	if checkFileIsExist(filePath) {
		f, err = os.OpenFile(filePath, os.O_WRONLY, 0777)
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filePath)
		fmt.Println("文件不存在")
	}
	_, err = io.WriteString(f, content)
	if err != nil {
		fmt.Println("写入失败", err)
		f.Close()
		return err
	} else {
		fmt.Println("写入成功")
		return nil
	}
}

// 读取 group 信息
func ReadGroupInfo() Groups {
	if !checkFileIsExist(confPath) {
		fmt.Println("group 配置不存在")
		return Groups{}
	}
	data, _ := ioutil.ReadFile(confPath)

	t := Groups{}
	yaml.Unmarshal(data, &t)
	fmt.Println("初始数据:", t)

	return t

	//d, _ := yaml.Marshal(&t)
	//fmt.Println(string(d))

	// 需要添加定时读取 yaml
}

func NewGroups() *GroupsG {
	newGroupsG := new(GroupsG)
	newGroupsG.GroupsG = ReadGroupInfo()
	return newGroupsG
}

// 对 yaml 增加一个 Group
func (groupsG *GroupsG) AddGroupInfo(jsonStr string) {
	group := &Group{}
	err := json.Unmarshal([]byte(jsonStr), &group)
	fmt.Println("group:", group)

	// 更新全局 groups
	groupsG.lock.Lock()
	groupsG.GroupsG = ReadGroupInfo()
	groupsG.lock.Unlock()

	var data []byte
	yaml.Unmarshal(data, &groupsG.GroupsG)

	groupsG.GroupsG.Group = append(groupsG.GroupsG.Group, *group)
	ret, err := yaml.Marshal(&groupsG.GroupsG)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ret))
		writeFile(string(ret), confPath)
	}
}

// 对 yaml 删除一个 Group
func (groupsG *GroupsG) DelGroupInfo(waitToDel string) {
	t := ReadGroupInfo()
	var newT Groups
	for _, v := range t.Group {
		if v.GroupName == waitToDel {
			continue
		}
		newT.Group = append(newT.Group, v)
	}

	// 更新全局 groups
	groupsG.lock.Lock()
	groupsG.GroupsG = newT
	groupsG.lock.Unlock()

	ret, err := yaml.Marshal(&newT)
	if err != nil {
		fmt.Println(err)
	} else {
		os.Rename(confPath, confPath + ".bak")
		fmt.Println("newT:", string(ret))
		err = writeFile(string(ret), confPath)
		if err != nil {
			// 写入失败，把 bak 文件恢复
			os.Remove(confPath)
			os.Rename(confPath + ".bak", confPath)
		} else {
			// 写入成功，删除 bak 文件
			os.Remove(confPath + ".bak")
		}
	}
}

func main() {
	testG := NewGroups()

	testG.AddGroupInfo("{\"GroupName\":\"groupTest2\", \"Previlege\":[{\"TagName\":\"tag2\", \"TagPrev\":\"admin\"}]}")
	testG.DelGroupInfo("groupTest1")

	var end chan struct{}
	end <- struct{}{}
}
