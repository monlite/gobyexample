// ticker是为了你想要定期重复做某事
package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker和timer的机制类似，它会向一个通道发送值。这里我们在通道上使用range来迭代通道中每500毫秒到达的一个值
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// ticker也可以被停止。一旦ticker停止了它就不会再从通道中收到任何值。我们在1600毫秒后停止它
	// 跑起这个程序，在我们停止ticker前它会tick3次
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
