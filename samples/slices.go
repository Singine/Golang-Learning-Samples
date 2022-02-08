package main

import (
	"fmt"
)

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:9]
	fmt.Printf("myslice为 %d, 其长度为: %d，长度为：%d\n", myslice, len(myslice),cap(myslice))
	fmt.Println(myslice)

	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[4])

	fmt.Println(myslice)
	fmt.Println(numbers4)

	newarr:=make([]int,cap(myslice))
	copy(newarr,myslice)
	fmt.Println(newarr)


	fmt.Printf("\n")
	fmt.Println(newarr)
	fmt.Printf("\n")

	newarr[3] = 666

	fmt.Println(newarr)
	fmt.Println(myslice)
	fmt.Println(numbers4)

}