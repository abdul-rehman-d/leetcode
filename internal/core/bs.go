package core

import "fmt"

// binary search
func BinarySearch(sortedArr []int, target int) int {
	start := 0
	end := len(sortedArr) - 1

	for start <= end {
		fmt.Println(sortedArr[start:end])
		middleIdx := start + ((end - start) / 2)
		middle := sortedArr[middleIdx]
		fmt.Println(middle)

		if middle == target {
			return middleIdx
		} else if middle > target {
			end = middleIdx - 1
		} else {
			start = middleIdx + 1
		}
	}

	return -1
}
