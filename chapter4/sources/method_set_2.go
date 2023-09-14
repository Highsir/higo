package main

import "./method_set_utils"

type Interface2 interface {
	M3()
	M4()
}

type T2 struct {}

func (t T) M3() {}
func (t *T) M4() {}

func main() {
	var t T
	var pt *T
	method_set_utils.DumpMethodSet(&t)
	method_set_utils.DumpMethodSet(&pt)
	method_set_utils.DumpMethodSet((*Interface2)(nil))
}

