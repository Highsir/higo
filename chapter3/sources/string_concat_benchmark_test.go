package chapter3

import "testing"
import "bytes"
import "strings"
import "fmt"


var s1 []string = []string{
	"Rob Pick ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByOperator(s1 []string) string {
	var s string
	for _, v := range s1 {
			s += v
		}
		return s
	}

func concatStringBySprintf(s1 []string) string {
	var s string
	for _, v := range s1 {
		s = fmt.Sprintf("%s%s", s,v)
        }
	return s
}

func concatStringByJoin(s1 []string) string{
	return strings.Join(s1, "")
}

func concatStringByStringsBuilder(s1 []string) string{
	var b strings.Builder
	for _, v := range s1 {
        b.WriteString(v)
    }
	return b.String()
}

func concatStringByStringsBuilderWithInitSize(s1 []string) string {
	var b strings.Builder
    b.Grow(len(s1))
    for _, v := range s1 {
        b.WriteString(v)
    }
    return b.String()
}

func concatStringByBytesBuffer(s1 []string) string {
	var b bytes.Buffer
	for _, v := range s1 {
        b.WriteString(v)
    }
	return b.String()
}

func concatStringByBytesBufferWithInitSize(s1 []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range s1 {
        b.WriteString(v)
    }
	return b.String()
}

func BenchmarkConcatStringbyOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByOperator(s1)
    }
}

func BenchmarkConcatStringbySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringBySprintf(s1)
    }
}

func BenchmarkConcatStringbyJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByJoin(s1)
    }
}

func BenchmarkConcatStringbyStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByStringsBuilder(s1)
    }
}

func BenchmarkConcatStringbyStringsBuilderWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByStringsBuilderWithInitSize(s1)
    }
}

func BenchmarkConcatStringbyBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByBytesBuffer(s1)
    }
}

func BenchmarkConcatStringbyBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
        concatStringByBytesBufferWithInitSize(s1)
    }
}


