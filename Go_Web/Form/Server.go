package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析url传递的参数，对于POST则解析响应包的主题request body
	// 如果没有ParseForm()则无法进行接下来的获取表单数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Println("vaal:", strings.Join(value, " "))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)    // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，执行登陆逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)        // 设置访问路由
	http.HandleFunc("/login", login)            // 设置访问路由
	err := http.ListenAndServe(":9091", nil)    // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}