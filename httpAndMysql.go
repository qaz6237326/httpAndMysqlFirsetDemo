package main

import (
	"net/http"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)



func helloHandler(w http.ResponseWriter, r *http.Request) {

	// 第一步
	// 连接mysql
	db, err := sql.Open("mysql", "root:123456@/userDB")
	// 所有事件执行完之后关闭链接
	defer db.Close()

	// 如果链接出错，则打印出错误
	if err != nil {
		panic(err)
	} else {
		// 如果成功了，打印出Success
		fmt.Println("Success")
	}

	// 第二步
	// 连接成功后，查询user表的信息
	rows, err := db.Query("SELECT name, age, createTime, updateTime FROM user")
	if err != nil {
		panic(err)
		return
	}

	// 第三步
	// 循环输出 查询到的信息
	for rows.Next() {

		//第四步
		// 先声明用来存放字段信息的变量
		var name string
		var createTime string
		var updateTime string
		var age int

		// 第五步
		// 扫描每一行的字段信息，把声明的变量地址作为参数进行匹配
		err := rows.Scan(&name, &age, &createTime, &updateTime)
		// 如何有错误的信息，则打印出来，不再继续执行下面的代码
		if err != nil {
			panic(err)
			return
		}

		// 第六步
		// 由于查询出来的时间格式需要进行转换才可以正确获取，所以下面要进行转换
		// UTC时区
		//newCreateTime, _ := time.Parse("2006-01-02 15:04:05", createTime)
		//newUpdateTime, _ := time.Parse("2006-01-02 15:04:05", updateTime)

		// 当前时区格式
		DefaultTime := time.Local
		newCreateTime, err := time.ParseInLocation("2006-01-02 15:04:05", createTime, DefaultTime)
		newUpdateTime, err := time.ParseInLocation("2006-01-02 15:04:05", updateTime, DefaultTime)

		 // 第七步打印
		fmt.Println(name, age, newCreateTime, newUpdateTime)
	}

	w.Write([]byte("Hello world!"))
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	listenErr := http.ListenAndServe(":9090", nil)
	if listenErr != nil {
		log.Fatal("listenAndServe: ", listenErr)
	}
	fmt.Println("服务启动成功")
}