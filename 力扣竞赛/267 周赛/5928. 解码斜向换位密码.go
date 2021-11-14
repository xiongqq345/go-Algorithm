package _67_周赛

import "strings"

//有 n 个人前来排队买票，其中第 0 人站在队伍 最前方 ，第 (n - 1) 人站在队伍 最后方 。
//
//给你一个下标从 0 开始的整数数组 tickets ，数组长度为 n ，其中第 i 人想要购买的票数为 tickets[i] 。
//
//每个人买票都需要用掉 恰好 1 秒 。一个人 一次只能买一张票 ，如果需要购买更多票，他必须走到  队尾 重新排队（瞬间 发生，不计时间）。如果一个人没有剩下需要买的票，那他将会 离开 队伍。
//
//返回位于位置 k（下标从 0 开始）的人完成买票需要的时间（以秒为单位）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/time-needed-to-buy-tickets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func decodeCiphertext(encodedText string, rows int) string {
	g := make([][]byte,rows)
	n := len(encodedText)
	for i := range g {
		g[i] = []byte(encodedText[n/rows*i : n/rows*(i+1)])
	}

	var i, j, col int
	var ans []byte
	for col < len(g[0]) {
		if inArea(g, i, j) {
			ans = append(ans, g[i][j])
			i, j = i+1, j+1
		} else {
			col++
			i, j = 0, col
		}
	}
	return strings.TrimRight(string(ans)," ")
}

func inArea(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}