// 在前面的示例中，我们使用与互斥锁的显式锁定来跨多个goroutine同步对共享状态的访问。
// 另一个选择是使用goroutines和channel的内置同步功能来实现相同的结果。
// 这种基于通道的方法与Go的通过沟通来共享内存的思想相符，这能保证在同一时间只有一个goroutine处理数据。
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中我们的状态将由一个goroutine拥有，这将保证数据永远不会因为并发访问而损坏，为了读取或修改状态，其他的goroutine将向拥有状态的goroutine发送消息并接收相应的回应。
// readOp和writeOp结构封装了这些请求，并且是拥有goroutine响应的一个方式
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 和以前一样，我们计算操作的次数
	var readOps uint64
	var writeOps uint64

	// reads和writes通道分别将被其他go协程用来发布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个是拥有state的goroutine，和前面例子中的map一样，不过这里是被这个状态协程私有的。
	// 这个goroutine在读取和写入通道上反复select，在请求到达时响应它们。
	// 先响应到达的请求，然后返回一个值到响应通道resp来表示操作成功（或者是reads中请求的值）。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动了100个goroutines，通过reads通道向状态goroutine发出读取请求。
	// 每次读取都需要构建一个readOp，通过reads通道发送它，并通过给定的resp通道接收结果
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法启动10个写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让go协程们跑1秒
	time.Sleep(time.Second)
	// 最后，获取并报告操作数
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
