package main

import "fmt"

type Test struct {
	Name string
}

func (test *Test)Get() (name string) {
	name = test.Name + "-Get"
	return
}

func (test *Test)Set(name string) {
	test.Name = name + "-Set"
	return
}

func main()  {
	t := Test{}
	t.Set("ok")
	fmt.Println(t.Get())
}