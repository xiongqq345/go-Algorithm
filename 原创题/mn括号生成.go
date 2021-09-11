package 原创题

// M对()，N对[],生成所有合法组合
func combinations(m, n int) []string {
	var ans []string
	var stack []int
	var helper func(string, int, int)
	helper = func(path string, leftM int, leftN int) {
		if leftM == 0 && leftN == 0 && len(stack) == 0 {
			ans = append(ans, path)
			return
		}

		if leftM > 0 {
			stack = append(stack, 1)
			helper(path+"(", leftM-1, leftN)
			stack = stack[:len(stack)-1]
		}
		if leftN > 0 {
			stack = append(stack, 2)
			helper(path+"[", leftM, leftN-1)
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			if stack[len(stack)-1] == 1 {
				stack = stack[:len(stack)-1]
				helper(path+")", leftM, leftN)
				stack = append(stack, 1)
			} else {
				stack = stack[:len(stack)-1]
				helper(path+"]", leftM, leftN)
				stack = append(stack, 2)
			}
		}
	}
	helper("", m, n)
	return ans
}
