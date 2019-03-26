package main

import "fmt"

// 这个person结构具有name和age域
type person struct {
	name string
	age  int
}

func main() {
	// 创建一个新的结构
	fmt.Println(person{"Bob", 20})

	// 也可以带上域的名字
	fmt.Println(person{name: "Alice", age: 30})

	// 忽略的域将被初始化成零值，这里是0
	fmt.Println(person{name: "Fred"})

	// &前缀初始化了一个指针指向结构
	fmt.Println(&person{name: "Ann", age: 40})

	// 通过.来访问结构中的域
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以使用.来访问结构指针中的域-指针被自动解引用了
	sp := &s
	fmt.Println(sp.age)

	// 结构是可变的
	sp.age = 51
	fmt.Println(sp.age)
}
