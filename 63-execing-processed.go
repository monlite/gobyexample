// 在前面的例子中，我们了解了生成外部进程的知识，当我们需要访问外部进程时时需要这样做。
// 但是有时候，我们只想用其他的（也许是非Go程序）来完全替代当前的Go进程。这时候，我们可以使用经典的exec方法的Go实现
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 在我们的例子中，我们将执行ls命令。Go需要提供我们需要执行的可执行文件的绝对路径，所以我们将使用exec.LookPath来得到它
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给ls一些基本的参数。注意，第一个参数需要是程序名
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec同样需要使用环境变量。这里我们仅提供当前的环境变量
	env := os.Environ()

	// 这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成/bin/ls -a -l -h进程。如果存在错误，那么我们将会得到一个返回值
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// go没有提供经典的Unix fork函数，这通常不是问题，因为使用go routine，生成进程，执行进程覆盖了大多数场景了
}
