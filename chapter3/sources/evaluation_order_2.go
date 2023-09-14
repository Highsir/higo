package main

import "fmt"

var (
    a = c + b
    b = f()
    _ = f()
    c = f()
    d = 3
)

func f() int {
    d++
    return d
}

func main() {
    fmt.Println(a, b, c, d)  // [a=10  _=5 c=6  b=4  d=6]
}