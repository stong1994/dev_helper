package ddl

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreateTableMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("abc")
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
		rst, err := NewCreateTableParser(content[0], CrateTableAdaptorAntlr{}).Parse()
		if err != nil {
			panic(err)
		}
		if err = t.Execute(w, rst); err != nil {
			panic(err)
		}
	}
}
