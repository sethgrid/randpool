package randpool

import (
	"testing"
)

func TestClock(t *testing.T) {
	c1 := NewClock(4)
	v1 := c1.Next()

	if v1 != 10001 {
		t.Errorf("got %d, want %d", v1, 10001)
	}

	c2 := NewClock(4)
	c2.Next()       // 10002
	c2.Next()       // 20002
	c2.Next()       // 30002
	v2 := c2.Next() // 40002

	if v2 != 40002 {
		t.Errorf("got %d, want %d", v2, 40002)
	}
}
