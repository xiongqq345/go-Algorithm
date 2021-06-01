package math

//给你一个下标从 0 开始的正整数数组candiesCount，其中candiesCount[i]表示你拥有的第i类糖果的数目。同时给你一个二维数组queries，其中queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]。
//
//你按照如下规则进行一场游戏：
//
//你从第0天开始吃糖果。
//你在吃完 所有第 i - 1类糖果之前，不能吃任何一颗第 i类糖果。
//在吃完所有糖果之前，你必须每天 至少吃 一颗糖果。
//请你构建一个布尔型数组answer，满足answer.length == queries.length 。answer[i]为true的条件是：在每天吃 不超过 dailyCapi颗糖果的前提下，你可以在第favoriteDayi天吃到第favoriteTypei类糖果；否则 answer[i]为 false。注意，只要满足上面 3 条规则中的第二条规则，你就可以在同一天吃不同类型的糖果。
//
//请你返回得到的数组answer。
func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make(map[int]int)
	sum[0] = candiesCount[0]
	res := make([]bool, len(queries))
	for i := 1; i < len(candiesCount); i++ {
		sum[i] = sum[i-1] + candiesCount[i]
	}
	for i, q := range queries {
		typ, l, r := q[0], q[1]+1, (q[1]+1)*q[2]
		res[i] = l <= sum[typ] && r > sum[typ-1]
	}
	return res
}
