package _1_双周赛

import "sort"

//一个整数数组 original 可以转变成一个 双倍 数组 changed ，转变方式为将 original 中每个元素 值乘以 2 加入数组中，然后将所有元素 随机打乱 。
//
//给你一个数组 changed ，如果 change 是 双倍 数组，那么请你返回 original数组，否则请返回空数组。original 的元素可以以 任意 顺序返回。
func findOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return nil
	}
	var ans []int
	sort.Ints(changed)
	var set [100005]int
	for _, v := range changed {
		set[v]++
	}

	for i := len(changed) - 1; i >= 0; i-- {
		if set[changed[i]] <= 0 {
			continue
		}
		if changed[i]%2 == 1 {
			return nil
		}
		if set[changed[i]/2] <= 0 {
			return nil
		}
		ans = append(ans, changed[i]/2)
		set[changed[i]/2]--
		set[changed[i]]--
	}
	return ans
}
