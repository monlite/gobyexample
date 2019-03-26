package main

import "fmt"

// 基本形式
func plus(a int, b int) int {
	return a + b
}

// 当有多个类型相同的形参时，可以省略前面的类型，只写最后一个
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res=plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
