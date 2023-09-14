package chapter3


import "fmt"

type Point struct {
    x float64
    y float64
}



func map_format_main() {
sl := []Point{
	{1.2345, 6.2789}, // Point{1.2345, 6.2789}
	{2.2345, 19.2789}, // Point{2.2345, 19.2789}
}

fmt.Println(sl)
// Go 1.5之前版本

m := map[Point]string{
    Point{29.935523, 52.891566}:   "Persepolis",
    Point{-25.352594, 131.034361}: "Uluru",
    Point{37.422455, -122.084306}: "Googleplex",
}
fmt.Println(m)

// Go 1.5及之后版本
m1 := map[Point]string{
    {29.935523, 52.891566}:   "Persepolis",
    {-25.352594, 131.034361}: "Uluru",
    {37.422455, -122.084306}: "Googleplex",
}
fmt.Println(m1)

m2 := map[string]Point{
    "Persepolis": {29.935523, 52.891566},
    "Uluru":      {-25.352594, 131.034361},
    "Googleplex": {37.422455, -122.084306},
}

fmt.Println(m2)
}
