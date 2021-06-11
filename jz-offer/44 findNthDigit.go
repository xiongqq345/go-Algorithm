package jz_offer

import (
	"math"
	"strconv"
)

// 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
//
//请写一个函数，求任意第n位对应的数字。
func findNthDigit(n int) int {
	var startNum, digit int
	for digit = 1; startNum < n; digit++ {
		next := digit * int(math.Pow10(digit)-math.Pow10(digit-1))
		if next > n {
			break
		}
		startNum += next
	}

	end := math.Pow10(digit-1) - 1 + math.Ceil(float64(n-startNum)/float64(digit))
	remainder := (n - startNum) % digit
	if remainder == 0 {
		return int(end) % 10
	}

	numStr := strconv.Itoa(int(end))
	return int(numStr[remainder-1] - 48)
}
