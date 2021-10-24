package math

//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
func majorityElement(nums []int) []int {
	var candidate1, candidate2 int
	var vote1, vote2 int
	for _, num := range nums {
		if vote1 > 0 && num == candidate1 {
			vote1++
		} else if vote2 > 0 && num == candidate2 {
			vote2++
		} else if vote1 == 0 {
			candidate1 = num
			vote1++
		} else if vote2 == 0 {
			candidate2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	var cnt1, cnt2 int
	for _, num := range nums {
		if vote1 > 0 && num == candidate1 {
			cnt1++
		}
		if vote2 > 0 && num == candidate2 {
			cnt2++
		}
	}
	var ans []int
	if cnt1 > len(nums)/3 {
		ans = append(ans, candidate1)
	}
	if cnt2 > len(nums)/3 {
		ans = append(ans, candidate2)
	}
	return ans
}
