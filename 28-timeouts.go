// 超时对于连接到外部资源或需要限制执行时间的程序非常重要。通过channel和select，在go中实现超时简单而优雅
package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设我们执行了一个外部调用，它会在2秒后把结果返回到c1中
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 这里select实现了超时机制。
	// res := <-c1等待调用结果， <- time.After将在等待1秒后超时返回
	// select将会选择最先准备好的case去执行，如果操作的时间超过1秒这里将会选择超时的case
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 如果允许长一点的超时时间，这里改成3秒，将会成功从c2中收到结果
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	// 运行程序将会显示第一次操作会超时，第二次会成功
	// select超时机制通过通道传递结果，这是个好主意
}
