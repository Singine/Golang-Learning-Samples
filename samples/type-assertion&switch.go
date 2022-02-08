
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = 10
	a := [...]string{}
	b := a[:]
	t1 := i.(int)
	fmt.Println(i)
	fmt.Println(t1)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

	fmt.Println("=====分隔线1=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, ok)

	fmt.Println("=====分隔线2=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d%t\n", t5, ok)

	fmt.Println("=====分隔线3=====")


	var tt interface{}
	var yy interface{} = 666
	var bb interface{} = "bobo"
	gg,hh := bb.(string)
	qq,ww := tt.(interface{})
	fmt.Println(gg,hh)
	fmt.Println(qq,ww)
	fmt.Println(reflect.TypeOf(tt))
	fmt.Println(reflect.TypeOf(yy))
	fmt.Println(reflect.TypeOf(bb))
}