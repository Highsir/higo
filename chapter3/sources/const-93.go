package chapter3

import "fmt"

const (
	a  = 5
	pi = 3.1415926
	s  = "Hello, Gopher"
	c  = 'a'
	b  = false
)

const (
	PI   = 3.1415926              // π
	PI_2 = 3.1415926 / (2 * iota) // π/2
	PI_4                          // π/4
)

type myInt int
type myFloat float32
type myString string

func const_main() {
	var j myInt = a
	var f myFloat = pi
	var str myString = s
	var e float64 = a + pi

	fmt.Println(j)               // 输出：5
	fmt.Println(f)               // 输出：3.1415926
	fmt.Println(str)             // 输出：Hello, Gopher
	fmt.Printf("%T, %v\n", e, e) // float64, 8.1415926

	fmt.Println(PI)   // 输出：3.1415926
	fmt.Println(PI_2) // 输出：3.1415926/2
	fmt.Println(PI_4) // 输出：3.1415926/4
}
