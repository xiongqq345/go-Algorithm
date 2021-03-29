package _74

import (
	"testing"

	"gotest.tools/assert"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {}}
	b := searchMatrix(matrix, 3)
	assert.Equal(t, b, true)
}
