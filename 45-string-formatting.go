// go为printf字符串格式化提供了出色的支持。以下是常见字符串格式化的一些示例
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	// go为常规go值的格式化提供了多种打印方式。例如，这里打印了point结构体的一个实例
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// 如果值是一个结构，%+v将包含结构的字段名
	fmt.Printf("%+v\n", p)

	// %#v将输出这个值的go语法表示，例如产生这个值的源码
	fmt.Printf("%#v\n", p)

	// 需要打印值的类型用%T
	fmt.Printf("%T\n", p)

	// 格式化布尔值
	fmt.Printf("%t\n", true)

	// 格式化整数有很多形式，使用%d进行标准的十进制格式化
	fmt.Printf("%d\n", 123)

	// 这个输出二进制的表示形式
	fmt.Printf("%b\n", 14)

	// 输出整数的字符表示形式
	fmt.Printf("%c\n", 33)

	// %x输出值的十六进制形式
	fmt.Printf("%x\n", 456)

	// 浮点数也有很多形式，使用%f进行标准的十进制格式化
	fmt.Printf("%f\n", 78.9)

	// %e和%E用科学计数法格式化
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 基本的字符串格式化用%s
	fmt.Printf("%s\n", "\"string\"")

	// 像go源码那样用双引号括起来输出用%q
	fmt.Printf("%q\n", "\"string\"")

	// 和上面的整形数一样，%x输出使用base-16编码的字符串，每个字节使用2个字符表示
	fmt.Printf("%x\n", "hex this")

	// %p输出指针的值
	fmt.Printf("%p\n", &p)

	// 格式化数字时，通常需要控制结果的宽度和精度。可以在%后面使用数字来控制宽度。默认情况下，结果将右对齐并用空格填充空白部分
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	//使用-进行左对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 也可以控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// 使用-进行左对齐
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// 到目前为止，我们已经看过Printf了，它通过os.Stdout输出格式化的字符串。Sprintf则格式化并返回一个字符串而不输出到任何地方
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 你可以使用Fprintf来格式化并输出到满足io.Writers接口的任何地方，而不是os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
