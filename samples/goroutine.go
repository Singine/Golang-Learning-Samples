package main

import (
	"fmt"
	"time"
)

func goroutineTest1() {
	fmt.Println("hello, go")
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
}

func goroutineTest2(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go goroutineTest1()
	fmt.Println("hello, world")

	go goroutineTest2("协程1号") // 第一个协程
	go goroutineTest2("协程2号") // 第二个协程
	time.Sleep(time.Second)
}