package _46

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lru(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	l.Put(3, 3)
	assert.Equal(t, l.Get(2), -1)
}
