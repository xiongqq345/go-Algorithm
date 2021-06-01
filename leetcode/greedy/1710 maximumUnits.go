package greedy

import "sort"

//请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：
//
//numberOfBoxesi 是类型 i 的箱子的数量。
//numberOfUnitsPerBoxi 是类型 i每个箱子可以装载的单元数量。
//整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。
//
//返回卡车可以装载单元 的 最大 总数。
func maximumUnits(boxTypes [][]int, truckSize int) (sum int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, v := range boxTypes {
		size := min(truckSize, v[0])
		sum += size * v[1]
		truckSize -= size
	}
	return sum
}
