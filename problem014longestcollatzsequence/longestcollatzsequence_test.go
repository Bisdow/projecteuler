package problem014longestcollatzsequence

import "testing"

func TestNewValue(t *testing.T) {
	if newValue(2) != 1 {
		t.Fail()
	}
	if newValue(3) != 10 {
		t.Fail()
	}
	if newValue(4) != 2 {
		t.Fail()
	}
	if newValue(5) != 16 {
		t.Fail()
	}
	if newValue(6) != 3 {
		t.Fail()
	}
	if newValue(7) != 22 {
		t.Fail()
	}
	if newValue(8) != 4 {
		t.Fail()
	}
	if newValue(9) != 28 {
		t.Fail()
	}
	if newValue(10) != 5 {
		t.Fail()
	}
}

func TestGetChainLength(t *testing.T) {
	if getChainLength(1) != 1 {
		t.Fail()
	}
	if getChainLength(2) != 2 {
		t.Fail()
	}
	if getChainLength(3) != 8 {
		t.Fail()
	}
	if getChainLength(4) != 3 {
		t.Fail()
	}
}
