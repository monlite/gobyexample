// 命令行标志是命令行程序指定选项的常用方式。例如，在wc -l中，这个-l就是一个命令行标志
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 基本的标记声明支持字符串、整数和布尔值选项，这里我们声明一个默认值为"foo"的字符串标志word并带有一个简短的描述。
	// 这里的flag.String函数返回一个字符串指针（不是一个字符串值），在下面我们会看到是如何使用这个指针的。
	wordPtr := flag.String("word", "foo", "a string")

	// 使用相同的方法来声明numb和fork标志
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。注意在标志声明函数中需要使用该参数的指针
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志都声明完成以后，调用flag.Parse来执行命令行解析
	flag.Parse()

	// 这里我们将仅输出解析的选项以及后面的位置参数。注意，我们需要使用类似*wordPtr这样的语法来对指针解引用，从而得到选项的实际值
	fmt.Println("word:", *wordPtr)
	fmt.Println("number:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)

	// 返回非标志参数，位置参数（位置参数要出现在标志参数后，否则，这个标志将会被解析为位置参数）
	fmt.Println("tail:", flag.Args())
}
