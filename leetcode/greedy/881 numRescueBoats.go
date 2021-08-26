package greedy

import "sort"

// 第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
//
//每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
//
//返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/boats-to-save-people
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1
	for i <= j {
		if people[j]+people[i] > limit {
			j--
		} else {
			i++
			j--
		}
	}
	return len(people) - j - 1
}
