package backtracking

//给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。
//
func addOperators(num string, target int) []string {
	n := len(num)
	var ans []string
	var helper func([]byte, int, int, int)
	helper = func(path []byte, idx, res, mul int) {
		if idx == len(num) {
			if res == target {
				ans = append(ans, string(path))
			}
			return
		}
		signIndex := len(path)
		if idx > 0 {
			path = append(path, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		for i, val := idx, 0; i < n && (i == idx || num[idx] != '0'); i++ {
			val = val*10 + int(num[i]-'0')
			path = append(path, num[i])
			if idx == 0 { // 表达式开头不能添加符号
				helper(path, i+1, val, val)
			} else { // 枚举符号
				path[signIndex] = '+'
				helper(path, i+1, res+val, val)
				path[signIndex] = '-'
				helper(path, i+1, res-val, -val)
				path[signIndex] = '*'
				helper(path, i+1, res-mul+mul*val, mul*val)
			}
		}

	}
	helper(make([]byte, 0, n*2-1), 0, 0, 0)
	return ans
}
