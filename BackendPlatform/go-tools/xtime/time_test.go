package xtime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration_UnmarshalText(t *testing.T) {
	var a = Duration(time.Second)
	b, err := json.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, "1000000000", string(b))
	var d Duration
	err = json.Unmarshal([]byte(`"1h"`), &d)
	assert.Nil(t, err)
	assert.Equal(t, Duration(3600000000000), d)
}
