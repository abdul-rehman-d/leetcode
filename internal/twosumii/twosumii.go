package twosumii

func TwoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) -1
	for {
		sum := numbers[lp] + numbers[rp]
		if target == sum {
			return []int{lp+1, rp+1}
		} else if target > sum {
			lp++
		} else {
			rp--
		}
	}
}

