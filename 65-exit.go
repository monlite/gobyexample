// 使用os.Exit来立即退出程序，并带有给定的状态码
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用了os.Exit时defer语句不会被运行，因此这个！永远不会输出
	defer fmt.Println("!")

	// 以状态码3退出
	os.Exit(3)

	// 注意，不像C语言在main中返回一个整数来指明退出状态。如果你想以非零状态退出，那么你就要使用os.Exit
}
