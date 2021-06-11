package jz_offer

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		sum1 := fmt.Sprintf("%d%d", nums[i], nums[j])
		sum2 := fmt.Sprintf("%d%d", nums[j], nums[i])
		return sum1 < sum2
	})
	var ans strings.Builder
	for _, num := range nums {
		ans.WriteString(strconv.Itoa(num))
	}
	return ans.String()
}
