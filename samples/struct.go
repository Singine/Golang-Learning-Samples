package main

import "fmt"

// 声明一个 Profile 的结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 重点在于这个星号: *
func (person *Profile) increase_age() {
	person.age += 1
}
func (person *Profile) change_name(v string) {
	person.name = v
}

func main() {
	myself := Profile{name: "小明", age: 24, gender: "male"}
	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increase_age()
	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.change_name("Gubao")
	fmt.Printf("name：%s", myself.name)

	xm := new(Profile)
	// 等价于: var xm *Profile = new(Profile)
	fmt.Println(xm)
	fmt.Println((*xm).age)
	// output: &{ 0 }

	xm.name = "Gubao"   // 或者 (*xm).name = "Gubao"
	xm.age = 3     //  或者 (*xm).age = 3
	xm.gender = "baby" // 或者 (*xm).gender = "baby"
	fmt.Println(xm)
	fmt.Println((*xm).name)


	bobo := Profile{name:"bobo"}
	gugu := &Profile{name:"gugu"}

	fmt.Println(bobo)
	fmt.Println(bobo.name)
	fmt.Println((*&bobo).name)
	fmt.Println(gugu.name)
	fmt.Println((&*gugu).name)



}