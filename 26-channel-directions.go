// 当使用channel作为函数参数的时候，你可以指定channel仅仅能够接收或者发送值，这样能够增加程序的类型安全性
package main

import "fmt"

// ping函数的参数pings只能够接收值，如果向它发送值将会编译错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong函数具有两个参数，一个pings发送值，一个pongs接收值
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
