// go支持指针，允许传递值的引用
package main

import "fmt"

// ival参数以值传递，zeroval会获取一个实参的副本
func zeroval(ival int) {
	ival = 0
}

// iptr是个指针，它保存的是实参的地址，在函数中对*iptr赋值，将会改变其指向地址的实际的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i将i的地址传递给函数
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 将打印出i的地址
	fmt.Println("pointer:", &i)
}
