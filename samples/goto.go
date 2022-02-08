package main

import "fmt"

func main() {
	//i := 1
//flag:
//	for i <= 10 {
//		if i%2 == 1 {
//			i++
//			goto flag
//		}
//		fmt.Println(i)
//		i++
//	}
	var j int = 1
GG:

	for j <= 10 {
		if j%2 == 1 {
			j++
			goto GG
		}
		fmt.Println(j)
		j++
	}

}
