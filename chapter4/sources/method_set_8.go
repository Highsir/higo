package main

// 如果结构体嵌入了多个接口类型的方法集合存在交集，
// 那么Go编译器将报错，除非结构体自己实现了交集中的所有方法

type Interface interface {
	M1()
	M2()
	M3()
}

type Interface1 interface {
	M1()
	M2()
	M4()
}

type T struct {
	Interface
	Interface1
}

func (T) M1() { println("T's M1") }
func (T) M2() { println("T's M2") }

func main() {
	t := T{}
	t.M1()
	t.M2()
}