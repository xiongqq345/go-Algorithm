package string

// 给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
//'A'：Absent，缺勤
//'L'：Late，迟到
//'P'：Present，到场
//如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//按 总出勤 计，学生缺勤（'A'）严格 少于两天。
//学生 不会 存在 连续 3 天或 3 天以上的迟到（'L'）记录。
//如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/student-attendance-record-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func checkRecord(s string) bool {
	var a, l int
	for _, v := range s {
		if v == 'A' {
			a++
			if a == 2 {
				return false
			}
		}
		if v == 'L' {
			l++
			if l == 3 {
				return false
			}
		} else {
			l = 0
		}
	}
	return true
}
