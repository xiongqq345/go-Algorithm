package sort

import (
	"fmt"
	"testing"
)

func Test_quick(t *testing.T) {
	arr := []int{5, 8, 1, 2, 9, 5, 1, 23, 4, 12, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
