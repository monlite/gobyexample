// go支持匿名函数，可以形成闭包，匿名函数在你想定义一个函数又不想命名的时候很有用
package main

import "fmt"

// intSeq这个函数返回了另外一个匿名的函数，返回的这个函数对i形成了闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	// 调用intSeq，将返回的结果（一个函数）赋值给nextInt
	nextInt := intSeq()

	// 闭包函数的效果，将会打印1，2,3
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 要确认该状态对于该特定函数是唯一的，创建并测试一个新函数
	newInts := intSeq()
	fmt.Println(newInts())
}
