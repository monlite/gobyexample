// 我们经常需要我们的程序对数据集合执行操作，例如选择满足给定条件的所有项，或使用自定义函数将所有项目映射到新集合
// 在某些语言中，会习惯使用泛型。go不支持泛型，在go中，当你的程序或者数据类型需要时，通常是通过组合的方式来提供操作函数
package main

import (
	"fmt"
	"strings"
)

// 这是一些string切片的组合函数示例。你可以使用这些例子来构建自己的函数。注意有时候，直接使用内联组合操作代码（匿名函数）会更清晰，而不是创建并调用一个帮助函数
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 返回目标字符串t出现的的第一个索引位置，或者在没有匹配值时返回-1
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果vs中的有一个字符串满足f函数，则返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果vs中所有的字符串都满足f函数，则返回true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个新的切片，包含vs中所有满足f函数的字符串
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个新的切片，包含将函数f应用于切片vs中每个字符串的结果
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	// 测试各种集合操作函数
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

	// 上面的例子都是用的匿名函数，但是你也可以使用类型正确的命名函数
}
