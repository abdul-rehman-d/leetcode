package koko_eating_bananas

import (
	"math"
	"slices"
)

func KokoEatingBananas(piles []int, h int) int {
	start := 1
	end := slices.Max(piles)

	ans := end
	for start < end {
		k := start + ((end - start) / 2)

		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}
		if hours > h {
			start = k + 1
		} else if hours == h {
			ans = min(k, ans)
			end = k
		} else {
			end = k
		}
	}

	return end
}
