package gid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostIP(t *testing.T) {
	ip, err := HostIP()
	t.Log(ip, err)
	assert.Nil(t, err)
	//assert.Equal(t, "192.168.27.61", ip.String())
}

func TestIpHashCode(t *testing.T) {
	code := IpHashCode() % (1 << 16)
	t.Log(code)
}
func TestHashCode(t *testing.T) {
	buf := []byte(`hello-world:formatString`)
	v1 := Fnv32a(buf)
	v2 := HashCode(buf)
	t.Log(v1)
	assert.Equal(t, v1, v2)
}

func TestSplitId(t *testing.T) {
	s := "5ba5eaae080b80b686725b38"
	low, high := SplitId(s)
	assert.Equal(t, low, "080b80b686725b38")
	v := StrToUint64(low)
	t.Log(v)
	t.Log(UnixFromStr(low))
	t.Log(FnvCodeFromStr(low))
	op := "hello-world:say-hello"
	vv := Fnv32aExt([]byte(op))
	opHash := fmt.Sprintf("%x", vv)
	t.Logf("%x", vv)
	assert.Equal(t, opHash, "5ba5eaae")
	assert.Equal(t, high, "5ba5eaae")

}

func TestHostnameHashCode(t *testing.T) {
	code :=HostnameHashCode()
	t.Log(code)
}
