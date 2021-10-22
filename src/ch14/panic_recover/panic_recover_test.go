package panic_recover_test

import (
	"errors"
	"fmt"
	_ "os"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
