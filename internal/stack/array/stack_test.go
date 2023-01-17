package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	s := NewStack()

	assert.Equal(t, s.Empty(), true, "Stack was not empty")
}

func TestNotEmpty(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	assert.Equal(t, s.Empty(), false, "Stack was empty")
}

func TestSizeZero(t *testing.T) {
	s := NewStack()
	assert.Equal(t, s.Size(), 0, "Expected zero elements found: %d", s.Size())
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	assert.Equal(t, s.Size(), 1, "Incorrect size")
}

func TestSizeThree(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	s.Add("Dave")
	s.Add("Sara")
	assert.Equal(t, s.Size(), 3, "Incorrect size")
}

func TestPopOne(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	v := s.Pop()
	assert.Equal(t, v, "Bob", "Expected Bob, found %s", v)
	assert.Equal(t, s.Size(), 0, "Incorrect size")
}

func TestPopTwo(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	s.Add("Alice")
	v := s.Pop()
	assert.Equal(t, v, "Alice", "Expected Alice, found %s", v)
	v = s.Pop()
	assert.Equal(t, v, "Bob", "Expected Bob, found %s", v)
}
