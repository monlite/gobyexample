// 写文件和读文件具有相似的模式
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 首先，把一个字符串写入文件
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 需要更精细的写入，先打开一个文件
	f, err := os.Create("/tmp/dat2")
	check(err)

	// go惯用在打开文件后立即用defer关闭文件
	defer f.Close()

	// 可以写入你期望的byte切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString同样可以用来写字符串
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用Sync来将写入的内容刷新到磁盘
	f.Sync()

	// bufio提供一个缓冲的writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用Flush确保所有缓冲的内容都写入到了底层的writer
	w.Flush()
}
