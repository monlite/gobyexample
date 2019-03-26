// go支持递归，这里是一个经典的例子：求阶乘
package main

import "fmt"

// fact函数调用自己直到它达到fact(0)的基本情况
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Print(fact(7))
}
