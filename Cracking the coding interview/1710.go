package coding

// 数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-majority-element-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func majorityElement(nums []int) int {
	var candidate, cnt int
	for _, v := range nums {
		if cnt == 0 {
			candidate = v
		}

		if v == candidate {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, v := range nums {
		if v == candidate {
			cnt++
		}
	}

	if cnt > len(nums)/2 {
		return candidate
	}
	return -1
}
