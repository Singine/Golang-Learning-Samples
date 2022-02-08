package main

import (
	"fmt"
	"reflect"
)

type Bobo struct {
	name string
	age int
}
type Baobao struct {
}

func (p Bobo)SayBye() string {
	return "Bye"
}
func (p Bobo)SayHello() string {
	return "Hello"
}
func (p Baobao)Baobao(name string, age int)  {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}

func main() {
	p := &Bobo{"Gubao", 27}
	q := Bobo{"Huihui", 27}
	z := Baobao{}
	testing_name := reflect.ValueOf("咕宝")
	testing_age := reflect.ValueOf(3)
	testing_arr := []reflect.Value{testing_name,testing_age}
	t1 := reflect.TypeOf(p)
	v1 := reflect.ValueOf(p)
	t2 := reflect.TypeOf(q)
	v2 := reflect.ValueOf(q)
	v3 := reflect.ValueOf(z)

	fmt.Println(t1)
	fmt.Println(v1)
	fmt.Println(v1.NumMethod())

	fmt.Println("--------------")

	fmt.Println(t2)
	fmt.Println(v2)
	fmt.Println(v2.NumMethod())

	fmt.Println("--------------")

	for i:=0;i<v1.NumMethod();i++{
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n", i+1,t1.Method(i).Name,v1.Elem().Method(i).Call(nil))
	}
	fmt.Println("--------------")
	v3.MethodByName("Baobao").Call(testing_arr)
}