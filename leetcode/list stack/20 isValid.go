package list_stack

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := range s {
		if _, ok := m[s[i]]; ok {
			if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
