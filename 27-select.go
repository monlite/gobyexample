// go的select可以让你等待多个通道操作，select结合goroutine和channel是go中强大的功能
package main

import (
	"fmt"
	"time"
)

func main() {
	// 在这个例子中我们将会在两个channel间select
	c1 := make(chan string)
	c2 := make(chan string)

	// 每个channel将会在一定时间后收到一个值，来模拟阻塞的RPC并发调用
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 我们用select同时等待两个值，并在到来时打印
	// 我们将先收到one，在收到two
	// 由于one和two的并发执行，总的用时是2s
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}
