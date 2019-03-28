// 关闭channel表示不会向它发送值了，这对于告诉channel接收者通信完成非常有用
package main

import "fmt"

func main() {
	// 在这个例子中，我们使用jobs通道来获取main goroutine中job被发送完的消息。当没有更多的job要发送给worker的时候，我们关闭jobs通道
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是worker goroutine，它不停的从jobs通道中取出job，在这里用两个值接收jobs的结果，
	// 当jobs通道被关闭并且通道里所有的值都已经接收了，more取出值将是false，这时我们通知done告诉main goroutine我们已经处理完所有job
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 这里发送3个job到jobs通道，然后关闭jobs通道
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 阻塞在这里等待worker线程完成
	<-done
}
