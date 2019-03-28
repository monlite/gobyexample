// 我们经常希望在未来的某个时刻执行Go代码，或者在某个时间间隔重复执行Go代码。go内置的timer和ticker能够非常简单的实现
package main

import (
	"fmt"
	"time"
)

func main() {
	// timer代表未来的一个事件，告诉timer你想等多久，它会提供一个在那时被通知的channel。这个timer会等2秒
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C会阻塞，直到timer发送一个值表示时间到了
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 如果你仅仅想等待，你可以使用time.Sleep。计时器很有用的一个原因是你可以在定时器到期前取消它，如下所示，提前调用Stop方法来停止定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
