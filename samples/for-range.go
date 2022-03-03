package main

import "fmt"

func main(){

	//for [condition |  ( init; condition; increment ) | Range]{
	//	statement(s);
	//}

	//1. 接一个条件表达式
	//2. 接三个表达式
	//3. 接一个 range 表达式 range后可接数组、切片，字符串等
	//4. 不接表达式 => 无限循环

	for i, l := range "go" {
		k := "dabei"
		fmt.Printf("%s,chishi\n",k)
		fmt.Println(i, l, string(l))
	}

	fmt.Println("")

	myarr := [...]string{"world", "python", "go"}
	for _,b := range myarr {
		fmt.Printf("hello, %v\n", b)
	}

}
