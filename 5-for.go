// for是go里唯一的循环语句
package main

import "fmt"

func main() {
	// 基本形式，单一的条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// 经典的for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 无限循环，直到break或者return才会退出循环
	for {
		fmt.Println("loop")
		break
	}

	// continue用来进入下一次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
