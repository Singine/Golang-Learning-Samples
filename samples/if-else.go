package main

import "fmt"

func agejudge(e int) {
	if age:= e; age > 18 {
		fmt.Println("已经成年了")
	} else if age >12 {
		fmt.Println("已经是青少年了")
	} else {
		fmt.Println("还不是青少年")
	}
}

func main() {
	agejudge(10)
}