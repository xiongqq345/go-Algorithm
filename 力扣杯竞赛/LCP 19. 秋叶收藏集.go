package 力扣杯竞赛

import "math"

//小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。
//出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。每部分树叶数量可以不相等，但均需大于等于 1。每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。
//
func minimumOperations(leaves string) int {
	const inf = math.MaxInt32
	n := len(leaves)
	f := make([][3]int, n)
	f[0][0] = boolToInt(leaves[0] == 'y')
	f[0][1] = inf
	f[0][2] = inf
	f[1][2] = inf
	for i := 1; i < n; i++ {
		isRed := boolToInt(leaves[i] == 'r')
		isYellow := boolToInt(leaves[i] == 'y')
		f[i][0] = f[i-1][0] + isYellow
		f[i][1] = min(f[i-1][0], f[i-1][1]) + isRed
		if i >= 2 {
			f[i][2] = min(f[i-1][1], f[i-1][2]) + isYellow
		}
	}
	return f[n-1][2]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
