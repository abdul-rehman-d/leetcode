package top_k_frequent_elements

func TopKFrequent(nums []int, k int) []int {

	res := make([][]int, len(nums))
	var out []int

	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}
	for num, count := range freq {
		res[count-1] = append(res[count-1], num)
	}

	var count int
	for i := len(res)-1; i>=0; i-- {
		if len(res[i]) == 0 {
			continue
		}
		out = append(out, res[i]...)
		count += len(res[i])
		if count == k {
			break
		}
	}

	return out
}

