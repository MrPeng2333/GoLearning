package flexible_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	fmt.Println("s1 == s2?", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	c3 := Customer{"2", "Mike", 40}
	fmt.Println("c1 == c2?", reflect.DeepEqual(c1, c2))
	fmt.Println("c1 == c3?", reflect.DeepEqual(c1, c3))

}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	fmt.Println(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	fmt.Println(*c)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
		return errors.New("the first param should be a pointer")
	}
	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
