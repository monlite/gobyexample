// 行过滤器是一种常见的程序，处理并打印结果到stdout，grep和sed都是常见的行过滤器
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用缓冲的scanner包装没有缓冲的os.Stdin，将获得一个方便的Scan方法，可以扫描scanner中的下一行
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text返回下一个读取的内容，这里是标准输入下一行
		ucl := strings.ToUpper(scanner.Text())
		// 输出大写的行
		fmt.Println(ucl)
	}

	// 检查扫描期间的错误，文件结束不会被当成错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
