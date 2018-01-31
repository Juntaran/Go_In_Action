/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/31 17:17
  */

package Blockchain

import (
	"fmt"
	"crypto/sha256"
	"crypto/rand"
)

type StringSet struct {
	set map[string]bool
}

func NewStringSet() StringSet {
	return StringSet{
		make(map[string]bool),
	}
}

func (set *StringSet) Add(str string) bool {
	_, found := set.set[str]
	set.set[str] = true
	return !found
}

func (set *StringSet) Keys() []string {
	var keys []string
	for k, _ := range set.set {
		keys = append(keys, k)
	}
	return keys
}

func CalHashSha256(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}

func PseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}