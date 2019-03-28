package main

import "fmt"

type rect struct {
	width, height int
}

// area方法的接收者是*rect类型
func (r *rect) area() int {
	return r.height * r.width
}

// 方法可以定义成指针或者值接收者，这是一个值接收者的例子
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// 指针调用时也用.即可，go会自动的解引用。当你想避免方法调用拷贝或者你想改变接收者的时候，用指针接收者
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
