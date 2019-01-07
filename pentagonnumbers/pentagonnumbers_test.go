package pentagonnumbers

import (
	"testing"
)

func TestCalcNextPenta(t *testing.T) {
	if 1 != lastElement() {
		t.Errorf("Not 1")
	}
	calcNextPenta()
	if 5 != lastElement() {
		t.Errorf("Not 5")
	}
	calcNextPenta()
	if 12 != lastElement() {
		t.Errorf("Not 12")
	}
	calcNextPenta()
	if 22 != lastElement() {
		t.Errorf("Not 22")
	}
	calcNextPenta()
	if 35 != lastElement() {
		t.Errorf("Not 35")
	}
	calcNextPenta()
	if 51 != lastElement() {
		t.Errorf("Not 51")
	}
	calcNextPenta()
	if 70 != lastElement() {
		t.Errorf("Not 70")
	}
	calcNextPenta()
}
