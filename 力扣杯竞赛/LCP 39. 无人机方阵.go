package 力扣杯竞赛

// 在 「力扣挑战赛」 开幕式的压轴节目 「无人机方阵」中，每一架无人机展示一种灯光颜色。 无人机方阵通过两种操作进行颜色图案变换：
//
//调整无人机的位置布局
//切换无人机展示的灯光颜色
//给定两个大小均为 N*M 的二维数组 source 和 target 表示无人机方阵表演的两种颜色图案，由于无人机切换灯光颜色的耗能很大，请返回从 source 到 target 最少需要多少架无人机切换灯光颜色。
func minimumSwitchingTimes(source [][]int, target [][]int) int {
	mp := make(map[int]int)
	for i := range target {
		for j := range target[i] {
			mp[target[i][j]]++
			mp[source[i][j]]--
		}
	}
	var ans int
	for _, v := range mp {
		if v > 0 {
			ans += v
		}
	}
	return ans
}
