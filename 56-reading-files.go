// 读写文件是许多程序的基本需求，先看下go读文件的例子
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读文件时需要检查很多调用错误，这里包装了个函数用来精简代码
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 最简单的读文件形式就是把文件内容全部读取到内存中
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 我们经常需要控制文件怎么被读，或者读哪一部分，对于这种任务，首先打开文件获取一个os.File
	f, err := os.Open("/tmp/dat")
	check(err)

	// 从文件的开头开始读一定的字节，这里最多允许读5个字节，不过也要注意我们实际读取了多少个（有可能比5个少）
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 你也可以用Seek定位文件的位置，并且从那个位置开始读
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// io包里提供了一些帮助方法，例如使用ReadAtLeast可以更稳健的实现上面的读取
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置退回的方法，不过可以用Seek(0, 0)来实现
	_, err = f.Seek(0, 0)
	check(err)

	// bufio包提供了缓冲读取，通常很有用。当有很多小而零碎的读取时能提升效率，并且它还提供了一些其他的读取方法
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 当读取完后用Close关闭文件（通常打开文件后立即使用defer来关闭文件）
	f.Close()
}
