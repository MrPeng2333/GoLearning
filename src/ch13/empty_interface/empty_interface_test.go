package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	switch i := p.(type) {
	case int:
		fmt.Println("Integer", i)
	case string:
		fmt.Println("string", i)
	default:
		fmt.Println("Unkown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
