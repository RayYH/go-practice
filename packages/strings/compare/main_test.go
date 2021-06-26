package main

import (
	"fmt"
	"strings"
	"testing"
)

func getA() string {
	return "a a a a a a a a a"
}

func getB() string {
	return "a a a a a a a a b"
}

// 10.68 ns/op
func BenchmarkCompare(b *testing.B) {
	b.ResetTimer()
	v1, v2, r := getA(), getB(), 0
	for i := 0; i < b.N; i++ {
		r = strings.Compare(v1, v2)
	}
	fmt.Print(r)
	b.StopTimer()
}

// 10.83 ns/op
func BenchmarkOperator(b *testing.B) {
	b.ResetTimer()
	v1, v2, r := getA(), getB(), 0
	for i := 0; i < b.N; i++ {
		if v1 == v2 {
			r = 0
		} else if v1 > v2 {
			r = 1
		} else {
			r = -1
		}
	}
	fmt.Print(r)
	b.StopTimer()
}
