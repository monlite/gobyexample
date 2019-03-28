// goroutine是轻量级的线程
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 通常的调用方式，同步执行
	f("direct")

	// 用go关键字调用，会新起一个goroutine调用函数。这个goroutine将和调用goroutine同步执行
	go f("goroutine")

	// 也可以用匿名函数的方式起goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 我们的两个函数调用现在在不同的goroutine中异步运行，因此主goroutine会到这里。这个Scanln需要我们在程序退出之前按一个键
	fmt.Scanln()
	fmt.Println("done")

	// 当我们运行这个程序时，我们首先看到阻塞调用的输出，然后是两个goroutine的交错输出。这种交错反映了Go运行时并发运行的goroutine
}
