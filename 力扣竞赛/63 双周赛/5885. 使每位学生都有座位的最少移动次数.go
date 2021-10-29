package _3_双周赛

import "sort"

//一个房间里有 n 个座位和 n 名学生，房间用一个数轴表示。给你一个长度为 n 的数组 seats ，其中 seats[i] 是第 i 个座位的位置。同时给你一个长度为 n 的数组 students ，其中 students[j] 是第 j 位学生的位置。
//
//你可以执行以下操作任意次：
//
//增加或者减少第 i 位学生的位置，每次变化量为 1 （也就是将第 i 位学生从位置 x 移动到 x + 1 或者 x - 1）
//请你返回使所有学生都有座位坐的 最少移动次数 ，并确保没有两位学生的座位相同。
//
//请注意，初始时有可能有多个座位或者多位学生在 同一 位置。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-number-of-moves-to-seat-everyone
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var ans int
	for i := range seats {
		ans += abs(students[i] - seats[i])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
