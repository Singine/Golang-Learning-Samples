package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func saybao(w http.ResponseWriter, r *http.Request)  {
	//解析模板
	t,err := template.ParseFiles("./template.tmpl")
	if err != nil{
		fmt.Println("Parse template failed, err:%v", err)
		return
	}

	//渲染模板
	str := "Gubao"
	err = t.Execute(w, str)
	if err!= nil{
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/",saybao)
	err := http.ListenAndServe(":5555",nil)
	if err!= nil{
		fmt.Println("Http Server failed, err:%v", err)
		return
	}
}
