package main

import "fmt"

func func01() {
	fmt.Println("在 func01 函数中，name：", name)
}

func main()  {
	var name string = "Python编程时光"
	fmt.Println("在 main 函数中，name：", name)

	func01()
}