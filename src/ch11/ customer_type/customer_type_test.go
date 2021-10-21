package customer_type

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntCov func(op int) int

func timeSpent(inner IntCov) IntCov {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}


func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
