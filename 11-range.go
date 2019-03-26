// range用来遍历各种数据结构中的元素
package main

import "fmt"

func main() {
	// range遍历切片，数组也一样
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 对数组和切片使用range，返回下标和值。上面例子中我们没有用到下标所以忽略，这个例子两个都用到
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 对map使用range来遍历键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 也可以仅仅对key遍历
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 用range遍历string中的unicode码点。第一个值返回的是下标，第二个是字符(rune)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
