package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name string
	Gender string
	Age int
}

func saybao(w http.ResponseWriter, r *http.Request)  {
	//解析模板
	t,err := template.ParseFiles("./template.tmpl")
	if err != nil{
		fmt.Println("Parse template failed, err:%v", err)
		return
	}

	//渲染模板
	//u1 :=UserInfo{
	//	Name: "Gubao",
	//	Gender: "Baby",
	//	Age: 4,
	//}

	m1 := map[string]interface{}{
		"Name": "Gubao",
		"Gender": "Baby",
		"Age": 5,
	}

	err = t.Execute(w,  m1)
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
