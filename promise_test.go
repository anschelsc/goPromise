package promise

import (
	"testing"
	"time"
)

func TestSanity(t *testing.T) {
	if 0 != 0 {
		t.Error("0 is not equal to 0, so the universe is meaningless.")
	}
	t1 := time.Nanoseconds()
	time.Sleep(5e8)
	t2 := time.Nanoseconds()
	if t1 == t2 {
		t.Error("Time didn't change during sleep.")
	}
	if t1 > t2 {
		t.Error("Time went backwards during sleep.")
	}
}

func TestEvaluation(t *testing.T) {
	p1 := Delay(time.Nanoseconds)
	p2 := Delay(time.Nanoseconds)
	p2V := p2.Force()[0].(int64)
	time.Sleep(5e8)
	p1V := p1.Force()[0].(int64)
	if p1V == p2V {
		t.Error("Promises conflated.")
	}
	if p1V < p2V {
		t.Error("Functions were evaluated at Delay time rather than Force time.")
	}
	if p2.Force()[0].(int64) != p2V {
		t.Error("Value not memoized.")
	}
}
