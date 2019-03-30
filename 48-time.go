// go为时间和周期提供了广泛支持，这里是一些例子
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 我们首先获取当前时间
	now := time.Now()
	p(now)

	// 你可以通过提供年，月，日等信息来创建一个time结构，这个结构总是和位置相关联，即时区
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 可以提取time结构中的各个时间相关的属性
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 打印是星期几
	p(then.Weekday())

	// 这些方法可以比较两个时间，测试是否比参数提供的提前、延后或相等
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub方法返回一个Duration，表示两个时间的间隔
	diff := now.Sub(then)
	p(diff)

	// 可以用不同的单位计算间隔的时间
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 使用Add来增加时间，或者使用-来倒退时间
	p(then.Add(diff))
	p(then.Add(-diff))
}
