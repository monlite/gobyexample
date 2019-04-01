// 环境变量是一个为Unix程序传递配置信息的普遍方式。让我们来看看如何设置，获取并列举环境变量
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用os.Setenv来设置一个键值对。使用os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// 使用os.Environ来获取环境中所有的键值对，返回一个key=value形式的切片，你可以用Split来分别获取key和value，这里我们打印所有的key
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

}
