/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 02:36
  */

package main

import (
	"fmt"
)


// map在函数传递映射过程中不会床在该映射的副本，对映射进行修改，所有这个映射的引用都会被修改

func removeColor(colors map[string]string, key string)  {
	delete(colors, key)
}

func main()  {
	// 创建一个映射，存储颜色和颜色对应的十六进制代码
	colors := map[string]string {
		"AliceBlue":	"#f0f8ff",
		"Coral":		"#ff7F50",
		"DarkGray":		"#a9a9a9",
		"ForestGreen":	"#228b22",
	}

	// 显示映射所有的颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
	// 调用函数移除指定的键
	removeColor(colors, "Coral")

	fmt.Println()
	// 显示映射所有的颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
