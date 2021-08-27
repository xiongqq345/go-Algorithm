package jz_offerII

//给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
//
//函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0 开始计数 ，所以答案数组应当满足 0 <= answer[0] < answer[1] < numbers.length 。
//
//假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kLl5u1
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			return []int{i, j}
		}
	}
	return nil
}
