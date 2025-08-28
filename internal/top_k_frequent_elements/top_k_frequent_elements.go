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
	for i := len(res) - 1; i >= 0; i-- {
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
func TopKFrequent2(nums []int, k int) []int {
	out := make([]int, k)
	freq := map[int]int{}     // num:len
	bucket := map[int][]int{} // len:num
	j := 0

	for _, num := range nums {
		freq[num]++
	}

	for num, value := range freq {
		bucket[value] = append(bucket[value], num)
	}

outer:
	for i := len(nums) + 1; i >= 0 && j != k; i-- {
		if v, ok := bucket[i]; ok {
			for x := range len(v) {
				out[j] = v[x]
				j++
				if j == k {
					break outer
				}
			}
		}
	}

	return out
}
