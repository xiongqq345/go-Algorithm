package jz_offerII

//给定一个整数数组 asteroids，表示在同一行的小行星。
//
//对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
//找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/XagZNi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func asteroidCollision(asteroids []int) []int {
	var st []int
	for _, v := range asteroids {
		f := true
		for len(st) > 0 && v < 0 && -v >= st[len(st)-1] && st[len(st)-1] > 0 {
			if -v == st[len(st)-1] {
				st = st[:len(st)-1]
				f = false
				break
			}
			st = st[:len(st)-1]
		}
		if f && (len(st) == 0 || (st[len(st)-1] < 0 && v < 0) || v > 0) {
			st = append(st, v)
		}
	}
	return st
}
