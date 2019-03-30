// 把数字转成字符串形式是基本且常见的需求，让我们看下在go中如何做
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 解析浮点数，这里的64表示解析多少位
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 对于ParseInt，0表示从推断的基数，64表示结果用64位去填充
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt也会识别16进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 也有无符号形式的ParseUint
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi可以方便的将字符串转换成10进制
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 当输入有误的时候，会产生一个错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
