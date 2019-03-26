// switch语句可以处理多分支条件的情况
package main

import (
	"fmt"
	"time"
)

func main() {
	// 基础的switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个case中可以用多个表达式，用逗号将它们隔开，这里我们也用了可选的default，表示默认情况
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 没有表达式的switch是if/else的另一种替代。这里也表示了case表达式可以是非常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型switch比较类型而不是值，可以用这种方式来判断interface的类型。这个例子中，将打印变量t对应的类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't konw type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
