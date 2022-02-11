package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayhi(w http.ResponseWriter, r *http.Request){
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}


func main() {
	http.HandleFunc("/gubao",sayhi)
	err :=http.ListenAndServe(":9090",nil)
	if err != nil{
		fmt.Println("http service error: %v\n",err)
		return
	}
}
