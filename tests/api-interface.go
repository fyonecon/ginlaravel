package main

import "fmt"

// ------------------------------

// Function 封装接口，仅暴露参数
type Function struct {
	Name string
}

func (fun *Function)Get() (name string) {
	name = fun.Name + "-Get"
	return
}

func (fun *Function)Set(name string)  {
	fun.Name = name + "-Set"
}

// ------------------------------

// Class 封装接口，仅暴露函数方法名
type Class interface {
	Set(string)
	Get() string
}

// ------------------------------

func main()  {

	// 直接调用Struct绑定的方法(函数)
	var f Function = Function{}
	f.Set("张三")
	fmt.Println(f.Get())

	// 调用Interface暴露了的方法(函数)
	var c Class = &Function{}
	c.Set("李四")
	fmt.Println(c.Get())

}