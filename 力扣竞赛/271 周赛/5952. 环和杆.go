package _71_周赛

import "strconv"

//总计有 n 个环，环的颜色可以是红、绿、蓝中的一种。这些环分布穿在 10 根编号为 0 到 9 的杆上。
//
//给你一个长度为 2n 的字符串 rings ，表示这 n 个环在杆上的分布。rings 中每两个字符形成一个 颜色位置对 ，用于描述每个环：
//
//第 i 对中的 第一个 字符表示第 i 个环的 颜色（'R'、'G'、'B'）。
//第 i 对中的 第二个 字符表示第 i 个环的 位置，也就是位于哪根杆上（'0' 到 '9'）。
//例如，"R3G2B1" 表示：共有 n == 3 个环，红色的环在编号为 3 的杆上，绿色的环在编号为 2 的杆上，蓝色的环在编号为 1 的杆上。
//
//找出所有集齐 全部三种颜色 环的杆，并返回这种杆的数量。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rings-and-rods
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func countPoints(rings string) int {
	n := len(rings) / 2
	var ans int
	var ring [10]map[byte]bool
	for i := 0; i < n; i++ {
		ring[i] = make(map[byte]bool)
	}
	for i := 0; i < len(rings)-2; {
		color := i
		i++
		for rings[i] != 'R' && rings[i] != 'G' && rings[i] != 'B' && i < len(rings) {
			i++
		}
		pos, _ := strconv.Atoi(rings[color+1 : i])
		ring[pos][rings[color]] = true
	}
	for i := range ring {
		if len(ring[i]) == 3 {
			ans++
		}
	}
	return ans
}
