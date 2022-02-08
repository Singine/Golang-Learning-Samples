package main
import "fmt"

type Student struct {
	name string
	age int
}

func main() {
	// new 一个内建类型
	num := new(int)
	fmt.Println(num)
	fmt.Println(*num) //打印零值：0

	// new 一个自定义类型
	s := new(Student)
	s.name = "wangbm"
	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(*s)

	fmt.Println("-----------------")

	a := make([]int, 2, 10)

	fmt.Printf("%T",a)
	fmt.Println()


}
