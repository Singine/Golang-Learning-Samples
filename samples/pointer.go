package main

import "fmt"

func main(){

	astr := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)

	bobo := new(string)

	fmt.Println(*bobo)

}

func mytest(ptr *int)  {
	fmt.Println(*ptr)
}