package main

import (
	"fmt"
)

func main() {
	rs := []rune{
        0x4E2D,
        0x56FD,
        0x6B22,
        0x8FCE,
        0x60A8,
    }
	s := string(rs)
	fmt.Println(s)

	sl := []byte{
        0xE4, 0xB8, 0xAD,
        0xE5, 0x9B, 0xBD,
        0xE6, 0xAC, 0xA2,
        0xE8, 0xBF, 0x8E,
        0xE6, 0x82, 0xA8,
    }

	s = string(sl)
	fmt.Println(s)
}
