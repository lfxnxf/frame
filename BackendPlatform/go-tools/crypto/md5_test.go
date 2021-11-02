package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	a := Md5("hello")
	assert.Equal(t, a, "5d41402abc4b2a76b9719d911017c592")
}

func TestMd5B(t *testing.T) {
	a := Md5B([]byte(`hello`))
	assert.Equal(t, a, "5d41402abc4b2a76b9719d911017c592")
}
