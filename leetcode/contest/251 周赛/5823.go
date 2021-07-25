package _51_周赛

func getLucky(s string, k int) int {
	var val int64
	for _, v := range s {
		num := int64(v) - 'a' + 1
		if num >= 10 {
			val += num/10 + num%10
		} else {
			val += num
		}
	}

	var pre int64
	for k > 1 {
		pre = 0
		for val > 0 {
			pre += val % 10
			val /= 10
		}
		val = pre
		k--
	}
	return int(val)
}
