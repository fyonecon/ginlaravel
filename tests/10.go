package main

import (
	"fmt"
	"time"
)

func hello1(num ...int) {
	num[1] = 28
	time.Sleep(1*time.Second)
}

func hello(num ...int) []int {
	num[0] = 18
	return num
}

func main()  {
	i := []int{5, 6, 7}

	//
	fmt.Println(i[1])
	hello1(i...)
	fmt.Println(i[1])

	//
	go hello(i...)
	fmt.Println(i[0])
	time.Sleep(1 * time.Second)
	fmt.Println(i[0])

}
