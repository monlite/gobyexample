// go对json的编解码提供了内置支持，包括内置和自定义数据类型
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 我们将使用这两个结构来演示下面的自定义类型的编码和解码
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 首先看一下对基本类型的编码成json字符串
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这里是切片和map的实例，它们将分别被编码成json数组和json对象
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON包可以自动编码您的自定义数据类型。编码输出中仅包含导出的字段（大写开头），并且默认情况下将这些字段名称用作JSON键
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 可以在struct字段声明上使用tag来自定义编码的JSON键名称。看下response2的定义，其中使用了小写的tag名
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 现在我们看下把json字符串解码为go值，这是一种通用的数据结构的例子
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码映射中的值，我们需要将它们转换为适当的类型。例如，我们将num中的值转换为预期的float64类型
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的数据需要一系列转换
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将JSON解码为自定义数据类型。这样做的好处是可以为我们的程序增加额外的类型安全性，并且在访问解码数据时不需要类型断言
	str := `{"page":1,"fruits":["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面的例子中，我们总是把json序列化输出在标准输出上，我们也可以直接序列化json到os.Writer接口上，例如os.Stdout甚至HTTP响应中
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
