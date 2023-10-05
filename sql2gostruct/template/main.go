package main

import (
	"html/template"
	"log"
	"net/http"
	"sql2gostruct/ddl"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.ParseFiles("static/ddl_create_table.gtpl")
			if err != nil {
				panic(err)
			}
			if err := t.Execute(w, nil); err != nil {
				panic(err)
			}
		} else {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			content := r.Form["ddl"]
			if len(content) == 0 {
				w.Write([]byte("输入为空"))
				return
			}
			t, err := template.ParseFiles("static/show_go_struct.gtpl")
			if err != nil {
				panic(err)
			}

			rst, err := ddl.CreateTableMethod(content[0])
			if err != nil {
				panic(err)
			}
			if err = t.Execute(w, rst); err != nil {
				panic(err)
			}
		}

	}) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
