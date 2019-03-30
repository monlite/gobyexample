// go内置了base64的编解码支持
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 这是我们将要编解码的字符串
	data := "abc123!?$*&()'-=@~"

	// go支持标准和URL兼容的base64，这里是标准的
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，你可能需要检查下它如果你不知道源字符串是不是合法形式
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 这是URL兼容的base64编解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	// 使用标准和URL编码出不太相同的值（最后是+或-），但是它们都可以分别解码为一样的原始字符串
}
