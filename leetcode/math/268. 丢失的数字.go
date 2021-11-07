package math

func missingNumber(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	n := len(nums)
	return n*(n+1)/2 - sum
}
