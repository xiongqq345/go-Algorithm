package _90

import (
	"fmt"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
