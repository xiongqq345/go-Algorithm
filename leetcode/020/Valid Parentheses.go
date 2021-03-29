package leetcode

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	var stack []byte
	m := make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
