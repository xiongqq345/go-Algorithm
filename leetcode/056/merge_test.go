package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func Test_merge(t *testing.T) {
	nums := [][]int{[]int{1, 4}, []int{4, 5}}
	assert.Equal(t, merge(nums)[0][1], 5)
}
