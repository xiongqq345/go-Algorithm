package _6_双周赛

//给你一个下标从 0 开始的字符串 street 。street 中每个字符要么是表示房屋的 'H' ，要么是表示空位的 '.' 。
//
//你可以在 空位 放置水桶，从相邻的房屋收集雨水。位置在 i - 1 或者 i + 1 的水桶可以收集位置为 i 处房屋的雨水。一个水桶如果相邻两个位置都有房屋，那么它可以收集 两个 房屋的雨水。
//
//在确保 每个 房屋旁边都 至少 有一个水桶的前提下，请你返回需要的 最少 水桶数。如果无解请返回 -1 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-number-of-buckets-required-to-collect-rainwater-from-houses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minimumBuckets(street string) int {
	var ans int
	bucket := -2
	for i := range street {
		if street[i] == 'H' {
			if bucket == i-1 {
			} else if i+1 < len(street) && street[i+1] == '.' {
				bucket = i + 1
				ans++
			} else if i > 0 && street[i-1] == '.' { // 再考虑在左侧放水桶
				ans++
			} else {
				return -1
			}
		}

	}
	return ans
}
