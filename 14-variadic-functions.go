// 可变参数函数可以用任意多的参数调用，例如fmp.Println就是这样的函数
package main

import "fmt"

// 这个函数可以接收任意数量的int
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// 可以使用单独的参数，以通常的方式调用
	sum(1, 2)
	sum(1, 2, 3)

	// 如果多个参数在slice中，用func(slice...)来调用
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
