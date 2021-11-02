package cast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToIntE2(t *testing.T) {
	v, err := ToIntE("aa")
	assert.Equal(t, v, 0)
	assert.NotNil(t, err)
}
