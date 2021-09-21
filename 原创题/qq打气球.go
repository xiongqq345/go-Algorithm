package 原创题

//【气球游戏】小Q在进行射击气球的游戏，如果小Q在连续T枪中打爆了所有颜色的气球，将得到一只QQ公仔作为奖励。（每种颜色的气球至少被打爆一只）。这个游戏中有m种不同颜色的气球，编号1到m。小Q一共有n发子弹，然后连续开了n枪。小Q想知道在这n枪中，打爆所有颜色的气球最少用了连续几枪？
//输入描述：
//第一行两个空格间隔的整数数n，m。n<=1000000 m<=2000
//第二行一共n个空格间隔的整数，分别表示每一枪打中的气球的颜色,0表示没打中任何颜色的气球。
//输出描述：
//一个整数表示小Q打爆所有颜色气球用的最少枪数。如果小Q无法在这n枪打爆所有颜色的气球，则输出-1

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n)
	var t int
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		arr[i] = t
	}
	output := f(n, m, arr)
	fmt.Println(output)
}

func f(n, m int, arr []int) int {
	minn := n + 1
	tmp := make([]int, m+1)
	for i, j := 0, 0; j < n; {
		for !contains(tmp) {
			tmp[arr[j]]++
			j++
		}

		minn = min(minn, j-i)
		tmp[arr[i]]--
		i++
	}
	if minn == n+1 {
		return -1
	}
	return minn
}

func contains(a []int) bool {
	for _, v := range a[1:] {
		if v <= 0 {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
