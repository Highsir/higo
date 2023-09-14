package main

import (
	"fmt"
)

func Max(n int, m int, f func(y int)){
	if n > m{
		f(n)
	} else{
		f(m)
	}
}

func main() {
	Max(2,3, func(y int) {fmt.Printf("%d\n", y)})
}