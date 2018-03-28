/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/27 19:05
  */

package druid

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func DoPost(serviceName, jsonString string)  {
	//fmt.Println("Do Post", serviceName, jsonString)
	url := "http://localhost:8200/v1/post/test"
	// json
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonString)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	if resp.Status != "200 OK" {
		fmt.Println("Error: response Status", resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
