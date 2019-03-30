// 程序中一个常见的需求是获取自Unix纪元以来的秒数，毫秒数或纳秒数，即时间戳
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 用Unix和UnixNano获取自Unix纪元以来的秒和纳秒数
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意没有UnixMillis，所以想获取毫秒数，需要自己用纳秒数计算
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以把秒或纳秒的时间戳转化为相应的Time结构
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
