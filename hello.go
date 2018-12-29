package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintln(w, "Hello World")
}

func main() {
	// 注册路由
	http.HandleFunc("/hello", IndexHandler)

	// 监听端口
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
