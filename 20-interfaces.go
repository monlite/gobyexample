// interface是方法签名的命名集合
package main

import (
	"fmt"
	"math"
)

// 这是一个几何图形的基本接口
type geometry interface {
	area() float64
	perim() float64
}

// 我们将用rect和circle实现这个接口
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 要实现一个接口，我们只需要实现接口中所有的方法即可
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了geometry，那么我们就可以调用接口里的方法。这是一个通用的measure方法，它可以用来测量任何实现了geometry接口的几何图形
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{
		width:  3,
		height: 4,
	}

	c := circle{
		radius: 5,
	}

	// rect和circle都实现了geometry接口，所以都可以调用measure方法
	measure(r)
	measure(c)
}
