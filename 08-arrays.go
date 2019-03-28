// 在go中，数组是特定长度的元素的编号序列
package main

import "fmt"

func main() {
	// 创建一个可以容纳5个int元素的数组，元素类型和数组长度都是数组类型的一部分。默认用零值填充，int的零值是0
	var a [5]int
	fmt.Println("emp:", a)

	// 可以用array[index] = val来设置值，或者用array[index]来取值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置方法len返回数组长度
	fmt.Println("len:", len(a))

	// 声明并初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
