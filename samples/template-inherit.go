package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func templateInherit_gubao(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	index,err := template.ParseFiles("./template-inherit_base.tmpl","./template-inherit_gubao.tmpl")
	//渲染模板
	if err != nil{
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	msg := "Hi,Gubao页面"
	index.ExecuteTemplate(w, "template-inherit_gubao.tmpl", msg)

}

func templateInherit_home(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	home,err := template.ParseFiles("./template-inherit_base.tmpl","./template-inherit_home.tmpl")
	//渲染模板
	if err != nil{
		fmt.Println("Parse template failed, err:%v", err)
		return
	}

	msg := "Hi,home页面"
	home.ExecuteTemplate(w, "template-inherit_home.tmpl", msg)
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
