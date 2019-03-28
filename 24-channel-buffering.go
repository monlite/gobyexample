// 默认channel是无缓冲的，意味着只有当有相应的接收者准备接收数据的时候，发送者才能向channel发送数据。带缓冲的channel可以接收一定数量的值哪怕没有接收者去接收
package main

import "fmt"

func main() {
	// 这里我们创建了一个能够缓冲2个值的channel
	messages := make(chan string, 2)
	// 因为channel具有缓冲，我们可以直接向channel发送值，而不用管对方有没有接收者
	messages <- "buffered"
	messages <- "channel"

	// 接下来我们接收通道中的两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
