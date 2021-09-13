package 力扣杯竞赛

import "sort"

//「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
//给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
func maxmiumScore(cards []int, cnt int) int {
	o, e := make([]int, len(cards)+1), make([]int, len(cards)+1)
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	var oidx, eidx int
	for _, v := range cards {
		if v&1 == 1 {
			o[oidx] = v
			oidx++
		} else {
			e[eidx] = v
			eidx++
		}
	}
	if cnt == 1 {
		return e[0]
	}
	oidx, eidx = 0, 0
	num := cnt
	var ans int
	for num > 0 {
		if o[oidx] > e[eidx] {
			ans += o[oidx]
			oidx++
		} else {
			ans += e[eidx]
			eidx++
		}
		num--
	}
	if ans&1 == 0 {
		return ans
	}

	//还有奇数,还有偶数
	if o[oidx] > 0 && e[eidx] > 0 {
		lasto := o[oidx-1]
		//如果已经选了偶数
		if eidx > 0 {
			laste := e[eidx-1]
			if lasto+o[oidx] > laste+e[eidx] {
				return ans - laste + o[oidx]
			} else {
				return ans - lasto + e[eidx]
			}
		}

		//如果没选偶数
		return ans - lasto + e[eidx]
	}

	// 没有奇数了，还有偶数
	if o[oidx] == 0 && e[eidx] > 0 {
		return ans - o[oidx-1] + e[eidx]
	}
	// 没有偶数了，还有奇数
	if eidx > 0 && e[eidx] == 0 && o[oidx] > 0 {
		return ans - e[eidx-1] + o[oidx]
	}

	//都没有了,说明不成立
	return 0
}
