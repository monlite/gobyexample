// 在这个例子中我们将用goroutine和channel实现工作池
package main

import (
	"fmt"
	"time"
)

// 这是worker，我们将会并行运行几个worker。这些worker会从jobs通道取出job，开始工作，然后把完成的job发送到results通道。每个工作都sleep了1秒，表示工作需要时间去完成
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// 为了使用工作池，我们需要用jobs通道给worker发送工作，并且用results通道收集他们的工作成果
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了3个worker，worker刚开始会阻塞，因为还没有job
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送了5个job然后关闭job表示我们总共就5个工作
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后我们从results收集worker的工作结果
	for a := 1; a <= 5; a++ {
		<-results
	}

	// 运行程序显示，5个job被各个worker执行。总工作时间为5秒，但是这个程序仅仅用2秒左右就完成了，是因为有3个worker同时工作
}
