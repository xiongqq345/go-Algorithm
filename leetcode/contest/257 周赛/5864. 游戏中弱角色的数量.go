package _57_周赛

import "sort"

// 你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] = [attacki, defensei] 表示游戏中第 i 个角色的属性。
//
//如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。更正式地，如果认为角色 i 弱于 存在的另一个角色 j ，那么 attackj > attacki 且 defensej > defensei 。
//
//返回 弱角色 的数量。
func numberOfWeakCharacters(p [][]int) int {
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] != p[j][0] {
			return p[i][0] > p[j][0]
		}
		return p[i][1] < p[j][1]
	})
	var ans, maxDef int
	for _, v := range p {
		if v[1] < maxDef {
			ans++
		}
		maxDef = max(maxDef, v[1])
	}
	return ans
}
