// defer用来确保一个函数调用稍后在程序中执行，通常用来做一些清理工作。defer用在像其他语言中的ensure和finally用到的地方
package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	// 假设我们想创建一个文件，往里面写一些东西，写完后关闭这个文件。这里展示如何用defer做到。
	// 在通过createFile获取一个文件对象后，立即用defer调用closeFile。
	// 它会在main函数结束时执行，就是writeFile结束后。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
