// URL是通用的资源定位的方式，看下go里怎么解析URL
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// 我们将解析此示例URL
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析该URL到URL结构，确保没有错误产生
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 解析Scheme
	fmt.Println(u.Scheme)

	// User包括用户名和密码
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host包括域名和端口，用SplitHostPort可以提取出来
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 提取path和#后面的fragment
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 获取字符串形式的查询参数
	fmt.Println(u.RawQuery)

	// 也可以解析到map中，map是map[string][]string类型的
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// 这里我们获取属性为k第一个参数值
	fmt.Println(m["k"][0])
}
