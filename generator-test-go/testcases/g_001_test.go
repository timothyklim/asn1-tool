package testcases

import (
	"generated-code/g_001"
	"testing"
)

func Test_g001_1(t *testing.T) {
	var value g_001.MyInteger = 10

	if value != 10 {
		t.Fatal("Wrong")
	}

}

func Test_g001_2(t *testing.T) {
	var value g_001.ColorType

	value.SetNavyBlue()

	if !value.IsNavyBlue() {
		t.Fatal("Wrong:", value)
	}
}
