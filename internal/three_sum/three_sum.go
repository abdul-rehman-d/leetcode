package three_sum

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := [][]int{}

	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a-1]==nums[a] {
			continue
		}
		b := a+1
		c := len(nums) -1
		for b < c && a < b && c < len(nums) && b < len(nums) {
			sum := nums[a] +nums[b] + nums[c]
			if sum < 0 {
				b++
			} else if sum > 0 {
				c--
			} else {
				res = append(res, []int{
					nums[a],
					nums[b],
					nums[c],
				})
				b++
				for nums[b-1] == nums[b] && b < c {
					b++
				}
			}
		}
	}

	return res
}

