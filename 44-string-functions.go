// 标准库的strings包提供了很多有用的字符串相关的函数。这里是一些用来让你对这个包有个初步了解的例子
package main

import "fmt"

// 给strings一个别名
import s "strings"

// 我们给fmt.Println一个短名字的别名，我们随后将会经常用到
var p = fmt.Println

func main() {
	// 这是一些strings中的函数例子。注意他们都是包中的函数，不是字符串对象自身的方法，这意味着我们在调用时需要将要操作的字符串作为第一个参数传递
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()

	// 虽然不是strings包的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}
