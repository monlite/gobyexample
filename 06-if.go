package main

import "fmt"

func main() {
	// 基本
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 可以不带else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 可以在条件前声明变量，声明的变量在所有的分支中都可以用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// 注意：在go中不需要括号，但是需要大括号
}
