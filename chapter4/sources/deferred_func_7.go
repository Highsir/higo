package chapter4

import (
	"fmt"
)

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	fmt.Println("foo1 result:")
	foo1()
	fmt.Println("foo2 result:")
	foo2()
	fmt.Println("foo3 result:")
	foo3()
}
