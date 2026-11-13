package median_of_two_sorted_arrays

import "math"

func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	// smaller array is A
	var A, B []int
	if len(nums1) <= len(nums2) {
		A = nums1
		B = nums2
	} else {
		A = nums2
		B = nums1
	}
	total := len(nums1) + len(nums2)
	halfOfBothArrays := total / 2

	l, r := 0, len(A)-1 // l and r are for A

	var median1, median2 int

	for {
		// guessing first half of merged array
		countA := l + (r-l)/2
		countB := halfOfBothArrays - countA
		mA := countA - 1 // middleIdx of array A
		mB := countB - 1 // middleIdx of array B

		var firstHalfRightMostA, firstHalfRightMostB, secondHalfLeftMostA, secondHalfLeftMostB int
		if mA < 0 {
			firstHalfRightMostA = math.MinInt64
		} else {
			firstHalfRightMostA = A[mA]
		}
		if mB < 0 {
			firstHalfRightMostB = math.MinInt64
		} else {
			firstHalfRightMostB = B[mB]
		}
		if mA+1 >= len(A) || mA+1 < 0 {
			secondHalfLeftMostA = math.MaxInt64
		} else {
			secondHalfLeftMostA = A[mA+1]
		}
		if mB+1 >= len(B) || mB+1 < 0 {
			secondHalfLeftMostB = math.MaxInt64
		} else {
			secondHalfLeftMostB = B[mB+1]
		}

		// is this guess right
		if firstHalfRightMostA <= secondHalfLeftMostB && firstHalfRightMostB <= secondHalfLeftMostA {
			// guess is right, therefore
			median1 = max(firstHalfRightMostA, firstHalfRightMostB)
			median2 = min(secondHalfLeftMostA, secondHalfLeftMostB)
			if total%2 == 1 {
				// one element
				// if min(len(nums1), len(nums2)) == 0 {
				// 	return float64(median1)
				// }
				return float64(median2)
			}
			// two elements
			return (float64(median1) + float64(median2)) / 2

		} else if firstHalfRightMostA > secondHalfLeftMostB {
			r = countA - 1
		} else {
			l = countA + 1
		}
	}
}
