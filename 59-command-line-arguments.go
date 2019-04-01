// 命令行参数是为命令行程序添加执行参数通用的方式，例如go run hello.go使用run和hello.go作为go程序的命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args提供对原始命令行参数的访问，注意切片中第一个值是程序的路径，os.Args[1:]是程序的参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 可以用下标获取特定的参数
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
