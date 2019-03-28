// 我们也可以通过for range来迭代地从channel中取出的值
package main

import "fmt"

func main() {
	// 我们将从通道中迭代2个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)
	// range迭代从queue中收到的每个值，因为我们上面关闭了通道，在迭代完2个值后就结束了
	// 这个例子也说明了可以关闭一个非空但是有值待取出的channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
