package main

import (
	"fmt"
	"reflect"
)

type Testing struct {
	name string
	age int
	gender string
}

type People struct {
	name string
	age int
	gender string
}

func (p People)SayBye()  {
	fmt.Println("Bye")
}

func (p People)SayHello()  {
	fmt.Println("Hello",p.name)
}


func main() {
	m := Testing{
		name: "Gubao",
	}

	u := "abc"

	t := reflect.TypeOf(&m)

	fmt.Println(m)
	fmt.Println(&m)
	fmt.Println(u)
	fmt.Println(&u)




	fmt.Println("&m Type: ",t)
	fmt.Println("&m Kind: ",t.Kind())

	fmt.Println("m Type: ",t.Elem())
	fmt.Println("m Kind: ",t.Elem().Kind())



	fmt.Println("-------------------------")




	p := People{"写代码的明哥", 27, "male"}

	v := reflect.ValueOf(p)
	q := reflect.TypeOf(p)

	fmt.Println("字段数:", v.NumField())
	fmt.Println("第 1 个字段：", v.Field(0))
	fmt.Println("第 2 个字段：", v.Field(1))
	fmt.Println("第 3 个字段：", v.Field(2))

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i:=0;i<v.NumField();i++{
		fmt.Printf("第 %d 个字段：%v \n", i+1, v.Field(i))
	}

	fmt.Println("==========================")

	fmt.Println("方法数:", q.NumMethod())
	for i:=0;i<q.NumMethod();i++{
		fmt.Printf("第 %d 个方法名字：%v \n", i+1, q.Method(i).Name)
	}

	p.SayHello()
}