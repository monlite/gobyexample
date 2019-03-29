// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器状态。对于更复杂的状态，我们可以使用互斥锁安全地访问多个goroutine中的数据
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 在这个例子中，state是个map
	var state = make(map[int]int)
	// mutex将用来同步对state的访问
	var mutex = &sync.Mutex{}

	// 用来统计执行的读写操作数量
	var readOps uint64
	var writeOps uint64

	// 在这里，我们启动100个goroutine来执行针对状态的重复读取，每个goroutine每毫秒执行一次
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// 对于每次读取，我们随机选择一个key，Lock()互斥锁以确保对state的独占访问，读取所选key的值，Unlock()解锁互斥锁，并增加readOps计数
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// 每次读取休息1毫秒
				time.Sleep(time.Millisecond)

			}
		}()
	}

	// 启动了10个goroutine用来写，每个goroutine每隔1毫秒会随机选择一个key，然后更新state的值
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 等待1秒，执行上面的流程
	time.Sleep(time.Second)

	// 这里读取并打印读和写的数量
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// 打印state最后的状态
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
