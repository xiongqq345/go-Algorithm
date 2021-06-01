package greedy

//假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
//给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(flowerbed, 0)
	flowerbed = append([]int{0}, flowerbed...)
	for i := 1; i < len(flowerbed)-1; i++ {
		// 已栽种则直接跳过一格
		if flowerbed[i] == 1 {
			i++
			continue
		}
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			i++
			n--
		}
		// 达到目标提前退出
		if n == 0 {
			return true
		}
	}
	return n <= 0
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	space := 1
	for _, v := range flowerbed {
		if v == 0 {
			space++
		} else {
			n -= (space - 1) / 2
			space = 0
		}
	}
	n -= space / 2
	return n <= 0
}
