package chapter4

import "fmt"

type BinaryAdder interface {
	Add(int, int) int
}

type MyAddFunc func(int, int) int

func (f MyAddFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func main() {
	var i BinaryAdder = MyAddFunc(MyAdd)
	fmt.Println(i.Add(5, 6))
}
