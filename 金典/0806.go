package coding

// 在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
//(1) 每次只能移动一个盘子;
//(2) 盘子只能从柱子顶端滑出移到下一根柱子;
//(3) 盘子只能叠在比它大的盘子上。
//
//请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
//
//你需要原地修改栈。
//
func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}

	var move func(int, *[]int, *[]int, *[]int)
	move = func(num int, from, mid, to *[]int) {
		if num < 0 {
			return
		}
		if num == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}

		move(num-1, from, to, mid)
		move(1, from, mid, to)
		move(num-1, mid, from, to)
	}
	move(len(A), &A, &B, &C)
	return C
}
