package redis

import (
	"testing"
)

func TestManager(t *testing.T) {
	r1 := m.Get("haha1")
	r1.Set("set1", 1)

	//	r2 := m.Get("haha2")
	//	r2.Set("set1", 2)
}
