package 力扣杯竞赛

//小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：
//
//小扣从 x 号站点移动至 x + 1 号站点需要花费的时间为 inc；
//小扣从 x 号站点移动至 x - 1 号站点需要花费的时间为 dec。
//现有 m 辆公交车，编号为 0 到 m-1。小扣也可以通过搭乘编号为 i 的公交车，从 x 号站点移动至 jump[i]*x 号站点，耗时仅为 cost[i]。小扣可以搭乘任意编号的公交车且搭乘公交次数不限。
//
//假定小扣起始站点记作 0，秋日市集站点记作 target，请返回小扣抵达秋日市集最少需要花费多少时间。由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。
//
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	const mod = int(1e9) + 7
	memo := map[int]int{}
	// 返回从起点 0 到达 end 站点所需最小时间
	var help func(end int) int
	help = func(end int) int {
		if end == 0 {
			return 0
		}
		// 从 0 站坐公交会回到原点，是无用功，肯定要步行
		if end == 1 {
			return inc
		}
		if res, ok := memo[end]; ok {
			return res
		}
		res := end * inc // 先假设全靠步行
		// 最后一步尝试坐每一辆公交
		for i, v := range jump {
			x := end / v
			// 从 x 站点坐公交达到的站点
			dest := x * v
			if dest == end { // end 可以整除 v
				res = min(res, cost[i]+help(x))
			} else {
				// 即 end 不能整除 v
				res = min(res, cost[i]+help(x)+inc*(end-dest))
				// 尝试从 x+1 坐公交之后步行返回的方案
				dest = (x + 1) * v
				res = min(res, cost[i]+help(x+1)+dec*(dest-end))
			}
		}
		memo[end] = res
		return res
	}
	return help(target) % mod
}
