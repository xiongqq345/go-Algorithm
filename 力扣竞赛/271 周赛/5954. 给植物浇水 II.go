package _71_周赛

func minimumRefill(plants []int, a int, b int) int {
	var ans int
	ca, cb := a, b
	for i, j := 0, len(plants)-1; i < j; i, j = i+1, j-1 {
		ca -= plants[i]
		cb -= plants[j]
		if i+1 > j-1 {
			break
		}
		if i+1 == j-1 {
			if ca < plants[i+1] && cb < plants[i+1] {
				ans++
			}
			break
		}
		if cb < plants[j-1] {
			ans++
			cb = b
		}
		if ca < plants[i+1] {
			ans++
			ca = a
		}
	}
	return ans
}
