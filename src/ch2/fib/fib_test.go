package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a, b := 1, 1
	fmt.Println(a, " ")
	for i := 0; i < 5; i++ {
		fmt.Println(b, " ")
		tmp := a
		a = b
		b = tmp + a
	}
}
