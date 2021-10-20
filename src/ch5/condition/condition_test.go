package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	} else {
		t.Log(a)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("unknow")

		}
	}
}
