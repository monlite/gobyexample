package main

import "fmt"

func main() {
	// 用make来创建空的map
	m := make(map[string]int)

	// 用map[key]=val来设置值
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// 用map[key]来获取值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 内置函数len获取map的键值对个数
	fmt.Println("len:", len(m))

	// 内置的delete函数来删除map中的键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// map[key]返回的可选的第二个值表示该key是否存在map中，用这种方法来判断更好，而不要用第一个返回值为0或""来判断。这里我们不需要第一个值，所以用_忽略它
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
