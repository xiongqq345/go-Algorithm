package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func Test_removeElement(t *testing.T) {
	assert.Equal(t, 5, removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
