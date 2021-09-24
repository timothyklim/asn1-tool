package testcases

import (
	"generated-code/g_007"
	"testing"
)

func Test_g007_1(t *testing.T) {
	var value g_007.Fruit

	value.SetBanana()

	if !value.IsBanana() {
		t.Fatal("Wrong:", value)
	}

	value.SetLemon()

	if !value.IsLemon() {
		t.Fatal("Wrong:", value)
	}
}

func Test_g007_2(t *testing.T) {
	var value g_007.TrafficLight

	value.SetNearlyRed()

	if !value.IsNearlyRed() {
		t.Fatal("Wrong:", value)
	}
}
