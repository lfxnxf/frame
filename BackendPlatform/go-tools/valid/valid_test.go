package valid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"127.0.0.1", true},
		{"127.0.0.300", false},
		{"127.0.0.a", false},
		{"aa", false},
	}
	for _, test := range tests {
		v := IP(test.input)
		assert.Equal(t, test.want, v)
	}
}
