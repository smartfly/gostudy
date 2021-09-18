package number_of_boomerangs

func numberOfBoomerangs(points [][]int) (ans int) {
	stat := map[int]int{}
	for _, x := range points {
		for _, y := range points {
			dis := (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
			stat[dis]++
		}
		for k, m := range stat {
			ans += m * (m - 1)
			delete(stat, k)
		}
	}
	return ans
}
