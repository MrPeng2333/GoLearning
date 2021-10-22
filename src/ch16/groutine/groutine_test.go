package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) { // 在一个Groutine里运行
			fmt.Println(i)
		}(i)
		time.Sleep(time.Millisecond * 50)
	}

}
