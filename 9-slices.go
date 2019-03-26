// 切片是go中的关键数据类型，为序列提供了比数组更强大的接口
package main

import "fmt"

func main() {
	// 不像数组，切片的类型仅仅由它所包含的元素类型决定（而不是元素数量）
	// 创建一个具有长度的空切片，用内置的make
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以像数组一样取值和设置值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len返回切片长度
	fmt.Println("len:", len(s))

	// 除了基本操作外，内置的append函数返回一个包含一个或多个新值的切片。注意我们需要接收append返回值，因为append可能创建了一个新的切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片也可以被复制，这里我们创建了一个新的切片，它和s的长度相同，并且将s拷贝到c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片支持切片操作，语法为slice[low:high]。例如，以下切片操作将获得s[2]，s[3]，s[4]
	l := s[2:5]
	fmt.Println("slc1:", l)

	// 从开始到5
	l = s[:5]
	fmt.Println("slc2:", l)

	// 从2到最后
	l = s[2:]
	fmt.Println("slc3:", l)

	// 也可以在一行声明并初始化切片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 切片可以组合成多维的，内层的切片的长度可以不固定，不像多维数组那样
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
