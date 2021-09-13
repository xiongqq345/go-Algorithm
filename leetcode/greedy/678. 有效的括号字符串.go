package greedy

func checkValidString(s string) bool {
	minLeft, maxLeft := 0, 0
	for _, ch := range s {
		if ch == '(' {
			minLeft++
			maxLeft++
		} else if ch == ')' {
			minLeft = max(minLeft-1, 0)
			maxLeft--
			if maxLeft < 0 {
				return false
			}
		} else {
			minLeft = max(minLeft-1, 0)
			maxLeft++
		}
	}
	return minLeft == 0
}
