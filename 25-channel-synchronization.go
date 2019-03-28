// 我们可以用channel在goroutine之间实现同步，这是个通过阻塞接收来等待完成的实例
package main

import (
	"fmt"
	"time"
)

// 这个函数会在一个goroutine中运行。done channel用来通知另一个goroutine工作已完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 向channel发送值，通知对方我们已经完成工作
	done <- true
}
func main() {
	// 创建一个工作线程，把用来通知的channel传进去
	done := make(chan bool, 1)
	go worker(done)

	// 这里会阻塞，直到收到从工作线程中传过来的值
	// 如果去掉<-done，程序会在工作完成前就退出
	<-done
}
