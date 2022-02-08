package main

import "fmt"


func main(){
	var (
		name string = "咕宝"
		action string = "啵啵你"

	)
	var p = &name
	var ptr = new(int)

	*ptr = 10



	fmt.Println(name + action)
	fmt.Println(p)
	fmt.Println("ptr address: ", ptr)
	fmt.Println("ptr value: ", *p)
	fmt.Println(*ptr)


}

