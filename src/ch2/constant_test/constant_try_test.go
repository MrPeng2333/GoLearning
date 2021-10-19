package constant_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednsday
)

func TestConstant(t *testing.T) {
	t.Log(Tuesday)
	t.Log(Wednsday)
}
