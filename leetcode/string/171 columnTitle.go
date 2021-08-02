package string

// 给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
func titleToNumber(columnTitle string) int {
	var ans int
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i, multiple = i-1, multiple*26 {
		ans += int(columnTitle[i]-'A'+1) * multiple
	}
	return ans
}
