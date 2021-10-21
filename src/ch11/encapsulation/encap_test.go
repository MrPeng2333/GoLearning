package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e1 := Employee{"000", "Bob", 18}
	e2 := Employee{Id: "001", Name: "Mike", Age: 22}
	e3 := new(Employee)
	e3.Id = "002"
	e3.Name = "Rose"
	e3.Age = 26
	t.Log(e1)
	t.Log(e2)
	t.Log(e3)
	t.Logf("e1 is %T", e1)
	t.Logf("e3 is %T", e3)
}

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"000", "Bob", 20}
	t.Log(e.String())
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
}
