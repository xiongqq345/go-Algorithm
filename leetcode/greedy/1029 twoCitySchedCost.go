package greedy

import "sort"

//公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。
//
//返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。
func twoCitySchedCost(costs [][]int) (cost int) {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	n := len(costs) / 2
	for i := 0; i < n; i++ {
		cost += costs[i][0] + costs[n+i][1]
	}
	return
}
