// go的sort包实现了对内置类型和用户自定义类型的排序，先看下内置类型的排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序方法是正对内置数据类型的。这里是一个字符串的例子。注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个int排序的例子
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 我们也可以使用sort来检查一个序列是不是已经是排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
