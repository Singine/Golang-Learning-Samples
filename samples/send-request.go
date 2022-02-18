package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("GET", "http://ip-api.com/json/101.229.179.144?lang=zh-CN", nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
