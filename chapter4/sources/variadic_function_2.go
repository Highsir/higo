package main

import "fmt"

func dump(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	//cannot use s (variable of type []string) as []interface{} value in argument to dump
	/* dump函数的变长参数类型为“...interface{}”，因此匹配该形参的要么是interface{}类型的变量，要么为“t...”（t类型为[]interface{}）。
	在例子中给dump传入的实参为“s...”，但s的类型为[]string，并非[]interface{}，导致不匹配。这里要注意的是，虽然string类型变量可以直接赋值给interface{}类型变量，
	但是[]string类型变量并不能直接赋值给[]interface{}类型变量。*/
	// s := []string{"Tony", "John", "Jim"} 

	s := []interface{}{"Tony", "John", "Jim"} 
	dump(s...)
}