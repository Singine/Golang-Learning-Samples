package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func templateInherit_gubao(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	index,err := template.ParseFiles("./templates-inherit_base.tmpl","./templates-inherit_gubao.tmpl")
	//渲染模板
	if err != nil{
		fmt.Println("Parse templates failed, err:%v", err)
		return
	}
	msg := "Hi,Gubao页面"
	index.ExecuteTemplate(w, "templates-inherit_gubao.tmpl", msg)

}

func templateInherit_home(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	home,err := template.ParseFiles("./templates-inherit_base.tmpl","./templates-inherit_home.tmpl")
	//渲染模板
	if err != nil{
		fmt.Println("Parse templates failed, err:%v", err)
		return
	}

	msg := "Hi,home页面"
	home.ExecuteTemplate(w, "templates-inherit_home.tmpl", msg)
}

func main() {
	http.HandleFunc("/gubao",templateInherit_gubao)
	http.HandleFunc("/home",templateInherit_home)

	err := http.ListenAndServe(":5555",nil)
	if err != nil{
		fmt.Println("Http Server failed, err:%v", err)
		return
	}

}
