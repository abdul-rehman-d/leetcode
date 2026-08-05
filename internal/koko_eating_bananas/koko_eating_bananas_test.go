package koko_eating_bananas

import (
	"testing"
)

type TestCase struct {
	Input  []int
	Hours  int
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
		{[]int{312884470}, 312884469, 2},
	}
	for _, tc := range testCases {
		result := KokoEatingBananas(tc.Input, tc.Hours)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
