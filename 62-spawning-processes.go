// 有时候我们的go程序需要生成其他的非go进程
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// 我们将从一个简单的命令开始，没有参数或输入，仅打印一些输出到标准输出流。exec.Command创建一个表示这个外部程序的对象
	dateCmd := exec.Command("date")

	// Output是另一个帮助我们处理运行命令的常见情况的函数，它等待命令运行完成，并收集命令的输出。如果没有出错，dateOut将获取到date命令的运行结果
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 下面我们将看看一个稍复杂的例子，我们将从外部进程的stdin输入数据并从stdout收集结果
	grepCmd := exec.Command("grep", "hello")

	// 这里我们明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出的结果，最后等待程序运行结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 上面的例子中，我们忽略了错误检测，但是你可以使用if err != nil的方式来进行错误检查，我们也只收集StdoutPipe的结果，但是你可以使用相同的方法收集StderrPipe的结果
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 如果你想使用通过一个字符串生成一个完整的命令，那么你可以使用bash命令的-c选项
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
