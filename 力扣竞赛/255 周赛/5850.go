package _55_周赛

// 给你一个整数数组 nums ，返回数组中最大数和最小数的 最大公约数 。
//
//两个数的 最大公约数 是能够被两个数整除的最大正整数。
func findGCD(nums []int) int {
	minn, maxx := nums[0], nums[0]
	for _, num := range nums {
		minn = min(minn, num)
		maxx = max(maxx, num)
	}
	return gcd(minn, maxx)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
