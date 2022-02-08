package main

import "fmt"

func main(){
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "xxxx":
		fmt.Println("xxxx")
	case s != "world":
		fmt.Println("world")
	}
}
