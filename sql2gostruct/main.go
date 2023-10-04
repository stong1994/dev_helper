package main

import (
	"log"
	"net/http"
	"sql2gostruct/ddl"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", ddl.CreateTableMethod) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil)    // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
