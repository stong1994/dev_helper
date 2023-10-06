package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sql2gostruct/ddl"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/ddl_create_table", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		fmt.Println(string(bytes))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		type createDDL struct {
			DDL string `json:"ddl"`
		}
		var request createDDL
		if err = json.Unmarshal(bytes, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.DDL == "" {
			w.Write(nil)
			return
		}
		rst, err := ddl.CreateTableMethod(request.DDL)
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		//data, _ := json.Marshal(map[string]interface{}{"struct": rst})
		if _, err = w.Write([]byte(rst)); err != nil {
			panic(err)
		}
	})

	fmt.Println("listening :9094")
	err := http.ListenAndServe(":9094", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
