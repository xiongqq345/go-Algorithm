package _51_周赛

func maximumNumber(num string, change []int) string {
	arr := []byte(num)
	var flag *bool
	for i, v := range arr {
		if int(v-'0') == change[int(v-'0')] {
			continue
		}
		if (flag == nil || *flag) && int(v-'0') < change[int(v-'0')] {
			arr[i] = byte(change[int(v-'0')]) + '0'
			flag = new(bool)
			*flag = true
		} else {
			if flag != nil {
				*flag = false
			}
		}
	}
	return string(arr)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
