// go具备返回多个值的能力，这个特性经常使用在函数中，同时返回结果和错误
package main

import "fmt"

// (int, int)表示这个函数返回两个int
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 用两个值接收函数调用结果
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 忽略其中一个
	_, c := vals()
	fmt.Println(c)
}
