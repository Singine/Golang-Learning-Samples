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
	t,err := template.ParseFiles("./templates.tmpl")
	if err != nil{
		fmt.Println("Parse templates failed, err:%v", err)
		return
	}

	//渲染模板
	u1 :=UserInfo{
		Name: "Gubao",
		Gender: "Baby",
		Age: 4,
	}

	m1 := map[string]interface{}{
		"Name": "Gubao",
		"Gender": "Baby",
		"Age": 5,
	}

	bao := []string{
		"灰灰",
		"药药",
		"嘿嘿",
		"边边",
		"獭獭",
		"龟龟",

	}

	err = t.Execute(w,  map[string]interface{}{
		"u1": u1,
		"m1":m1,
		"bao":bao,
	})



	if err!= nil{
		fmt.Println("Parse templates failed, err:%v", err)
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
