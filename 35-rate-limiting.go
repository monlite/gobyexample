// 速率限制是控制资源利用和维持服务质量的重要机制。go可以通过goroutine，channnel和ticker优雅的实现
package main

import (
	"fmt"
	"time"
)

func main() {
	// 首先我们来看下基本的速率限制。假设我们想限制到来请求的处理速度。我们将通过requests通道提供这些请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 这个limiter通道每200毫秒会收到一个值。这是我们速率限制中的监管器
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在服务每个请求之前阻塞来自limiter通道的接收，我们将速率限制为每200毫秒1个请求
	for req := range requests {
		<-limiter
		fmt.Println("request:", req, time.Now())
	}

	// 我们可能希望在我们的速率限制方案中允许短时间的突发请求，同时保留总体速率限制。
	// 我们可以通过缓冲的limiter通道来实现这一目标。此burstyLimiter通道将允许最多3个事件的突发。
	burstyLimiter := make(chan time.Time, 3)

	// 填充3个请求到通道中，表示允许的突发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每200毫秒我们会添加一个值到burstyLimiter，最多能达到它的最大容量：3个
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟5个请求的到来
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 其中前3个将会突发请求，受益于burstyLimiter，可以看到前3个请求的时间几乎没有间隔。后来2个请求间隔200毫秒
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request:", req, time.Now())
	}
}
