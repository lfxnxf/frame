package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64Join(t *testing.T) {
	a := []int64{1, 1, 3, 4}
	assert.Equal(t, Int64Join(a, ","), "1,1,3,4")
	a = []int64{}
	assert.Equal(t, Int64Join(a, ","), "")
}

func TestRmDup(t *testing.T) {
	a := []string{"a", "a", "c","c"}
	b := RmDup(a)
	assert.Equal(t, 2, len(b))
	a = []string{}
	b = RmDup(a)
	assert.Equal(t, 0, len(b))
}
