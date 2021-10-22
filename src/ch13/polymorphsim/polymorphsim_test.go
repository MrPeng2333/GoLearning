package polymorphsim_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProGrammer struct{}

func (p *JavaProGrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphsim(t *testing.T) {
	var prog Programmer = &GoProgrammer{}
	// goProg := new(GoProgrammer)
	javaProg := new(JavaProGrammer)
	writeFirstProgram(prog)
	writeFirstProgram(javaProg)
}
