package implement_queue_using_stacks

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	value := obj.Peek()
	assert.Equal(t, 1, value)
	obj.Push(2)
	value = obj.Peek()
	assert.Equal(t, 1, value)
	value = obj.Pop()
	assert.Equal(t, 1, value)
	value = obj.Peek()
	assert.Equal(t, 2, value)

	assert.False(t, obj.Empty())

	value = obj.Pop()
	assert.Equal(t, 2, value)

	assert.True(t, obj.Empty())
}
