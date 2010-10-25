package promise

import (
	"testing"
)

func TestSanity(t *testing.T) {
	if 0 != 0 {
		t.Errorf("0 is not equal to 0, so the universe is meaningless")
	}
}

func TestEvaluation(t *testing.T) {
	f1 := func() interface{} {
		return 3
	}
	p := Delay(f1)
	if p.value != nil {
		t.Errorf("p.value should be %s, but it was %s", nil, p.value)
	}
	if val := p.Force(); val != 3 {
		t.Errorf("p.Force should have returned %s, got %s", 3, val)
	}
	if p.value != 3 {
		t.Errorf("p.value should be %s, but it was %s", 3, p.value)
	}
}
