package main

import "fmt"

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	myname := myfunc() 								//调用函数myfunc 先打印了#12 再return name = go myname变量被赋值为return的值 最后运行defer name重新赋值为 python
	fmt.Printf("main 函数里的name: %s\n", name) 		//打印现在的name 是重新赋值的python
	fmt.Println("main 函数里的myname: ", myname) 	//打印myname变量的值 go
}


