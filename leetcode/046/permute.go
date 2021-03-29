package _46

func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	var path []int
	used := make([]bool, n)
	dfs(&res, path, nums, used, n, 0)
	return res
}

func dfs(res *[][]int, path, nums []int, used []bool, length, depth int) {
	if depth == length {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i, v := range nums {
		if used[i] {
			continue
		}
		path = append(path, v)
		used[i] = true
		dfs(res, path, nums, used, length, depth+1)
		path = path[:len(path)-1]
		used[i] = false
	}
}
