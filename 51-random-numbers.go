// go的math/rand包提供了随机数生成
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Intn返回一个随机数n，0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64返回一个float64随机数，0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// 也可以生成其他范围的随机数。这里5.0 <= f' < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 默认数字生成器是确定性的，因此默认情况下每次都会产生相同的数字序列。为了产生不同的序列，给它一个变化的随机种子，这里用当前的时间当种子。
	// 请注意，如果你想要保密的随机数，这是不安全的，请使用crypto/rand或其他方式产生随机数
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果用固定的数当随机种子，则会生成固定的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
