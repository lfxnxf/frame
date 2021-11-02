package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Offer(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, q.Len(), 0)
	q.Offer(1)
	q.Offer(2)
	assert.Equal(t, q.Len(), 2)
	q.Poll()
	assert.Equal(t, q.Len(), 1)
	v := q.Poll()
	vv := v.(int)
	assert.Equal(t, vv, 2)
	assert.True(t, q.Empty())
	n := q.Poll()
	nn, ok := n.(int)
	t.Log(nn, ok)
	assert.Nil(t, n)
	assert.False(t, ok)

}

func TestNewStack(t *testing.T) {
	s := NewStack()
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.Empty(), true)
	s.Push(1)
	s.Push(2)
	assert.Equal(t, s.Len(), 2)
	v0, _ := s.Peek().(int)
	assert.Equal(t, v0, 2)
	v1, _ := s.Pop().(int)
	assert.Equal(t, v1, 2)
	v2, _ := s.Pop().(int)
	assert.Equal(t, v2, 1)
	v3, ok := s.Pop().(int)
	assert.Equal(t, v3, 0)
	assert.Equal(t, ok, false)

}
