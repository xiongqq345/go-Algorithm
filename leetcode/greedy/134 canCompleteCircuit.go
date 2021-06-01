package greedy

//在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。你从其中的一个加油站出发，开始时油箱为空。
//
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
//
//说明:
//
//如果题目有解，该答案即为唯一答案。
//输入数组均为非空数组，且长度相同。
//输入数组中的元素均为非负数。
func canCompleteCircuit(gas []int, cost []int) int {
	var low, ans, cur int
	for i := range gas {
		cur += gas[i] - cost[i]
		if cur < low {
			low = cur
			ans = i + 1
		}
	}
	if cur < 0 {
		return -1
	}
	return ans
}
