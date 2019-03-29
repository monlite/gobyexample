// Go中管理状态的主要机制是通过通道进行通信。我们在工作池中看到了这个例子。但是，还有一些其他方式可用于管理状态。在这里，我们将看一下sync/atomic包用于多个goroutines访问的原子计数器
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// 我们使用一个无符号整数来表示计数器
	var ops uint64

	// 用50个goroutine来模拟并发更新计数器，每个goroutine每隔1毫秒对计数器进行加1操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 为了以原子方式递增计数器，我们使用AddUint64，把ops计数器的地址传给它
				atomic.AddUint64(&ops, 1)
				// 每次增加的过程中休息1毫秒
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 等待1秒，让计数器累加
	time.Sleep(time.Second)
	// 为了在其他goroutine仍在更新时安全地使用计数器，我们通过LoadUint64将当前值的副本提取到opsFinal中。如上所述，我们需要为此函数提供计数器ops的内存地址
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
