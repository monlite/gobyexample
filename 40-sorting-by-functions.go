// 有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
// 例如，我们想按照字母的长度而不是首字母顺序对字符串排序。这里是一个go自定义排序的例子。
package main

import (
	"fmt"
	"sort"
)

// 为了在go中使用自定义函数进行排序，我们需要一个对应的类型。这里我们创建一个为内置[]string类型的别名的byLength类型
type byLength []string

// 我们在类型中实现了sort.Interface的Len，Less和Swap方法，这样我们就可以使用sort包的通用Sort方法了。
// Len和Swap通常在各个类型中都差不多，Less将控制实际的自定义排序逻辑。
// 在我们的例子中，我们想按字符串长度增加的顺序来排序，所以这里使用了len(s[i])和len(s[j])。
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// 一切都准备好了，我们现在可以通过将原始的fruits切片转型成byLength来实现我们的自定排序了。
	fruits := []string{"peach", "banana", "kiwi"}
	// 对这个转型的切片使用sort.Sort方法。
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// 类似的，参照这个我们可以创建一个自定义类型，实现这个类型的这三个接口方法，然后在一个这个自定义类型的集合上调用sort.Sort方法，我们就可以使用任意的函数来排序go切片了
}
