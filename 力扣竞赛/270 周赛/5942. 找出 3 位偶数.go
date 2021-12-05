package _70_周赛

import "sort"

//给你一个整数数组 digits ，其中每个元素是一个数字（0 - 9）。数组中可能存在重复元素。
//
//你需要找出 所有 满足下述条件且 互不相同 的整数：
//
//该整数由 digits 中的三个元素按 任意 顺序 依次连接 组成。
//该整数不含 前导零
//该整数是一个 偶数
//例如，给定的 digits 是 [1, 2, 3] ，整数 132 和 312 满足上面列出的全部条件。
//
//将找出的所有互不相同的整数按 递增顺序 排列，并以数组形式返回。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/finding-3-digit-even-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findEvenNumbers(nums []int) []int {
	n := len(nums)
	var list, ans []int
	vis := make([]bool, n)
	var helper func(int, int)
	helper = func(val, num int) {
		if num == 3 {
			list = append(list, val)
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && nums[i-1] == nums[i] && vis[i-1] {
				continue
			}
			if vis[i] {
				continue
			}
			if val == 0 && nums[i] == 0 {
				continue
			}
			if num == 2 && nums[i]%2 == 1 {
				continue
			}
			vis[i] = true
			helper(10*val+nums[i], num+1)
			vis[i] = false
		}
	}
	helper(0, 0)

	mp := make(map[int]bool)
	for _, v := range list {
		if !mp[v] {
			mp[v] = true
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return ans
}

func findEvenNumbers(digits []int) (ans []int) {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}
next:
	for i := 100; i < 1000; i += 2 { // 枚举所有三位数偶数 i
		c := [10]int{}
		for x := i; x > 0; x /= 10 { // 枚举 i 的每一位 d
			d := x % 10
			if c[d]++; c[d] > cnt[d] { // 如果 d 出现次数比 digits 中的 d 的次数还多，那么 i 肯定不能由 digits 中的数字组成
				continue next // 枚举下一个偶数
			}
		}
		ans = append(ans, i)
	}
	return
}
