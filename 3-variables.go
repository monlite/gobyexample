package main

import "fmt"

func main() {
	// var声明一个或多个变量
	var a = "initial"
	fmt.Println(a)

	// 可以一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动类型推导
	var d = true
	fmt.Println(d)

	// 没有初始值的声明，将被初始化为零值，int的零值是0
	var e int
	fmt.Println(e)

	// := 是声明并初始化的简写形式，只能用在局部作用于中
	f := "short"
	fmt.Println(f)
}
