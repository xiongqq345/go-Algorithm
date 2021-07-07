package binary_search

// 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用bool isBadVersion(version)接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

func firstBadVersion(n int) int {
	i, j := 1, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if isBadVersion(h) {
			j = h
		} else {
			i = h + 1
		}
	}
	return i
}

func isBadVersion(int) bool {
	return false
}
