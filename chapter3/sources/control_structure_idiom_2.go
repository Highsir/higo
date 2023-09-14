package main

import ("fmt")

func pointerToArrayRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("pointerToArrayRangeExpression result:")
    fmt.Println("a = ", a)

	for i,v := range &a {
		if i  == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt .Println("a = ", a)
}

func sliceRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("pointerToArrayRangeExpression result:")
    fmt.Println("a = ", a)

	for i,v := range a[:] {
		if i  == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt .Println("a = ", a)
}

// 切片与数组还有一个不同点，就是其len在运行时可以被改变，而数组的长度可认为是一个常量，不可改变。那么len变化的切片对for range有何影响呢？
func sliceLenChangeRangeExpression(){
	var a = []int{1, 2, 3, 4, 5}
    var r = make([]int, 0)

    fmt.Println("sliceLenChangeRangeExpression result:")
    fmt.Println("a = ", a)

	for i, v := range a {
        if i == 0 {
            a = append(a, 6, 7)
        }

        r = append(r, v)
    }

    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}



func main() {
	// pointerToArrayRangeExpression()
	// sliceRangeExpression()
	sliceLenChangeRangeExpression()

}