// channel是连接并发goroutine的管道。您可以将值从一个goroutine发送到通道，并从另一个goroutine中接收它
package main

import "fmt"

func main() {
	// 用make创建一个channel，这个channel的类型是string，代表它只能传递string类型的值
	message := make(chan string)
	// 通过channel<-来向通道发送值。这里我们从一个新的goroutine中发送了一个ping到channel中
	go func() {
		message <- "ping"
	}()

	// 通过<-channel语法来从通道中接收值。这里我们接收上面发送的ping并且把它打印出来
	// 默认情况下，发送和接收端会阻塞直到他们都准备就绪（channel的缓冲是0）。这种特性允许我们在程序结束时等待ping消息的到达，而不用使用其他的同步方式
	msg := <-message
	fmt.Println(msg)
}
