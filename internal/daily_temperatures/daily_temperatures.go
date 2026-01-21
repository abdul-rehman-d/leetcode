package daily_temperatures

func DailyTemperatures(temperatures []int) []int {
	out := make([]int, len(temperatures))

	hm := make(map[int]int)
	for i := len(temperatures) - 1; i >= 0; i-- {
		t := temperatures[i]
		hm[t] = i

		for x := 100; x > t; x-- {
			if idx, ok := hm[x]; ok && idx > i {
				if out[i] == 0 {
					out[i] = idx - i
				} else {
					out[i] = min(idx-i, out[i])
				}
			}
		}
	}

	return out
}
