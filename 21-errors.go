// go传递错误的惯用方式是：通过单独的明确返回值。这与java和ruby等语言中使用异常和C使用单个值表示错误码不同。go的方式更容易看到哪些方法产生错误，并且处理错误和处理非错误时的语言结构一样。
package main

import (
	"errors"
	"fmt"
)

// 按照惯例，错误在返回列表的最后，并且是一个error类型，它是一个内置的接口类型
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New用给定的错误消息构造一个基本的错误
		return -1, errors.New("can't work with 42")
	}
	// 在error的位置返回nil表示没有错误
	return arg + 3, nil
}

// 可以自己来实现error接口来创建新的错误类型，只要实现Error方法就行了。这里是上面例子的一个变体，使用自定义的错误类型
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 使用&argError构造一个新的结构
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 下面对两种错误返回的方式进行了测试。注意在if中使用inline形式的错误检查是go的习惯用法
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果你想使用自定义错误中的数据，你需要通过类型断言来获取自定义错误的实例
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
