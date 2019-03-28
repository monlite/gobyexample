// 默认情况下向channel发送或接收是阻塞的。然而，我们可以使用带default子句的select来实现非阻塞发送，接收，甚至非阻塞多路选择
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这是非阻塞接收。如果有值可以从messages中接收，select会选择接收值，如果没有的话，会立刻进入default分支
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// 类似地，这是非阻塞发送。这里msg不能发送因为messages没有缓冲，并且没有接收者就绪，因此会进入default分支
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在带default的情况下，可以有多个case，用来实现非阻塞多路选择，这里我们同时从messages和signals非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}
