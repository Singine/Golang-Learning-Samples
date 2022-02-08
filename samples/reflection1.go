package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)
	fmt.Println(t)
	fmt.Println(v)

	var kk interface{} = "666"
	fmt.Println(kk,reflect.TypeOf(kk))

	fmt.Printf("%T \n", kk)
	fmt.Printf("%T \n", reflect.TypeOf(kk))

	kkk,ok := reflect.ValueOf(kk).Interface().(string)


	fmt.Printf("222222222%T \n", kkk)

	fmt.Println(kkk,ok)
	fmt.Println("-------------------------")

	name := "Go编程时光"

	tt := reflect.ValueOf(name)
	qq := reflect.ValueOf(&name)

	fmt.Println(tt)
	fmt.Println(qq)
	fmt.Println("-------------------------")
	fmt.Println("可写性为:", tt.CanSet())
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(*&name)
	fmt.Println(reflect.TypeOf(&name))
	fmt.Println(qq)
	fmt.Println(qq,reflect.TypeOf(qq))
	fmt.Println(qq.Elem(),reflect.TypeOf(qq.Elem()))
	fmt.Println(qq.Elem().CanSet())
	fmt.Println(qq.CanSet())






}