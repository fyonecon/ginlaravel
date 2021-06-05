package main

import (
	"fmt"
)

func main()  {

	a1 := []int{0, 1, 2, 3, 4, 5}
	//a1 = append(a1, 61)
	//a1 = append(a1, 62)
	//a1 = append(a1, 63)
	b1 := append(a1, 2021)
	//b1 = append(b1, 2022)
	//b1 = append(b1, 2023)
	//b1 = append(b1, 2024)
	a1 = append(a1, 64)
	a1 = append(a1, 65)
	//a1 = append(a1, 66)

	fmt.Println(a1)
	fmt.Println(b1)

	//a1 := []int{0, 1, 2}
	//a1 = append(a1, 1)
	//b1 := append(a1, 2021)
	//b1 = append(b1, 2022)
	//b1 = append(b1, 2023)
	//b1 = append(b1, 2024)
	//b1 = append(b1, 2025)
	//b1 = append(b1, 2026)
	//b1 = append(b1, 2027)
	//b1 = append(b1, 2028)
	//b1 = append(b1, 2029)
	//
	//fmt.Println(a1)
	//fmt.Println(b1)
	//
	//fmt.Printf("%p %p \n", a1, b1)
	//fmt.Println(&a1, &b1, cap(a1), cap(b1))

	//a2 := [...]int{0, 1, 2, 3}
	//_b2 := make([]int, 0)
	//_a2 := a2[0:len(a2)]
	//b2 := append(_b2, _a2...)
	//b2 = append(b2, 2021)
	//fmt.Println(b2)

	//a1 := []int{0, 1, 2, 3}
	//a2 := [...]int{10, 11, 12, 13}
	//b := make([]int, 0)
	//
	//c := append(b, 2021)
	//d := append(a1, 2022)
	//e := append(b, a2[0:len(a2)]...)
	//
	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)

}
