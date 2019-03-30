// SHA1哈希经常用于计算二进制或文本blob的短标识。例如，git版本控制系统广泛使用SHA1来识别版本化文件和目录。以下是在Go中计算SHA1哈希的方法
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// 生成sha1哈希首先要New
	h := sha1.New()

	// 写入想要生成的对象
	h.Write([]byte(s))

	// Sum获取最终hash的结果到一个字符切片，它的参数用来追加到一个已存在的[]byte上，但是通常是不需要的
	bs := h.Sum(nil)

	// sha1值通常用16进制打印
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// 如果你想计算MD5哈希的话，使用crypto/md5包中的md5.New
}
