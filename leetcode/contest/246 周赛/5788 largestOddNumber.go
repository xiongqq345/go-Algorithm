package _46_周赛

func largestOddNumber(num string) string {
	i, j := 0, len(num)-1
	for ; j >= 0; j-- {
		if (num[j]-48)%2 == 1 {
			break
		}
	}
	if j < 0 {
		return ""
	}
	for ; i <= j; i++ {
		if num[i] != '0' {
			break
		}
	}
	return num[i : j+1]
}
