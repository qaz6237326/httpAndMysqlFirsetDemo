package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"encoding/json"
)

// 定义结构体
type Student struct {
	Name string
	Age  int
	Sex  string
	Colors []string
}




func main() {
	// 注册路由
	http.HandleFunc("/hello", IndexHandler)

	// 监听端口
	err2 := http.ListenAndServe("127.0.0.1:8000", nil)
	if err2 != nil {
		log.Fatal("ListenAndServe: ", err2)
	}
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 声明一个变量，获取结构体，并且赋值
	var stu = Student{"dtc", 22, "boy", []string{"red", "blue", "green"}}

	// 转换成json格式的数据流
	j, err := json.Marshal(stu)
	if err != nil {
		fmt.Errorf("Marshal Error %v", err)
	}
	// 转换成用string()解析流打印出来
	fmt.Println(string(j))
	r.ParseForm()
	fmt.Println(r.Form) //得到的数据集合 map[key:[value],key2:[value],...]
	fmt.Println("path", r.URL.Path) // 访问的路径
	fmt.Println("scheme", r.URL.Scheme) // scheme
	fmt.Println(r.Form["url_long"]) // []

	// 循环输出数据
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 把hello world 写入 w 里，返回给前端
	// fmt.Fprintln(w, "Hello World")
	fmt.Fprintln(w, string(j))
}