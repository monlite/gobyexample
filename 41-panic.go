// panic意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
package main

import "os"

func main() {
	// panic通常用法就是在一个函数返回了错误值但是我们不知道如何处理（或不想处理）时终止运行。这里是一个在创建一个新文件时返回异常错误时的panic用法。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	// 注意，不像在有些语言中使用异常处理错误，在go中则习惯通过返回值来表示错误
}
