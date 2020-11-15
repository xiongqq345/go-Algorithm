package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func Test_threeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	assert.Equal(t, threeSumClosest(nums, 1), 2)
}
