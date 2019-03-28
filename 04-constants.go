// go支持字符、字符串、布尔、数字常量
package main

import (
	"fmt"
	"math"
)

// const声明了一个常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// const可以出现在var能够出现的任何地方
	const n = 500000000

	// 常量表达式可以以任意精度表示
	const d = 3e20 / n

	// 数字常量没有类型，除非它被给定一个，例如显式转换
	fmt.Println(int64(d))

	// 在需要的时候，会根据上下文给常量一个类型，例如变量赋值或者函数调用的时候。这里math.Sin需要一个float64
	fmt.Println(math.Sin(n))

}
