package math

// 已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
//
//不要使用系统的 Math.random() 方法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-rand10-using-rand7
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

func rand7() int {

}
