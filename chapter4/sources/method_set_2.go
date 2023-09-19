package main

type Interface2 interface {
	M3()
	M4()
}

type T2 struct{}

func (t T2) M3()  {}
func (t *T2) M4() {}

func main() {
	var t T2
	var pt *T2
	DumpMethodSet(&t)
	DumpMethodSet(&pt)
	DumpMethodSet((*Interface2)(nil))
}
