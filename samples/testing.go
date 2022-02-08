package main

import (
	"fmt"
	"time"
)


func couting(b bool)  {
	if b{
		for a:=10;a>=0;a--{
			fmt.Println(a)
			time.Sleep(time.Millisecond)
		}
	}else{
		for a:=0;a<=10;a++{
			fmt.Println(a)
			time.Sleep(time.Millisecond)
		}
	}
}


func main() {

	go couting(false)
	go couting(true)


	time.Sleep(time.Second)
}