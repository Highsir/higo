package chapter4

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTye := v.Elem()

	n := elemTye.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTye)
	}

	fmt.Printf("%s's method set:\n", elemTye)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTye.Method(j).Name)
	}
}
