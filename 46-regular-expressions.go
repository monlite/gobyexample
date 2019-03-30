// go提供内置的正则表达式，这是正则表达式的例子
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 这测试一个字符串是否符合正则表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们直接使用了字符串表示的正则表达式，但是想要把正则表达式应用于其他的匹配任务，最好把正则表达式编译到结构中，以便以后使用
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试
	fmt.Println(r.MatchString("peach"))

	// 这是查找从最左边开始第一个匹配到的字符串
	fmt.Println(r.FindString("peach punch"))

	// 这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容
	fmt.Println(r.FindStringIndex("peach punch"))

	// Submatch返回完全匹配和局部匹配的字符串。例如，这里会返回p([a-z]+)ch和([a-z]+)的信息
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 类似的，这个会返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带All的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 带All的也可以也可用于我们上面看到的其他功能
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 提供非负整数作为函数的第二个参数将限制匹配的数量
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的例子的参数是string类型，而且函数名字都带有String，像MatchString。你也可以使用[]byte类型参数，把函数名字中的String去掉就行了
	fmt.Println(r.Match([]byte("peach")))

	// 使用正则表达式创建常量时，可以使用Compile的MustCompile变体。普通的Compile不适用于常量，因为它有2个返回值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp包还可用于将字符串子集替换为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func变体允许您使用给定函数转换匹配的文本
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
