package math

//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//
//例如，
//
//    1 -> A
//    2 -> B
//    3 -> C
//    ...
//    26 -> Z
//    27 -> AA
//    28 -> AB
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/excel-sheet-column-title
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		remainder := (columnNumber-1)%26 + 1
		ans = append(ans, byte('A'+remainder-1))
		columnNumber = (columnNumber - remainder) / 26
	}
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}
