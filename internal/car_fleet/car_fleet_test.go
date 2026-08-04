package carfleet

import (
	"testing"
)

type I struct {
	Target   int
	Position []int
	Speed    []int
}

type TestCase struct {
	Input  I
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{
			I{Target: 12, Position: []int{10, 8, 0, 5, 3}, Speed: []int{2, 4, 1, 1, 3}},
			3,
		},
		{
			I{Target: 10, Position: []int{3}, Speed: []int{3}},
			1,
		},
		{
			I{Target: 100, Position: []int{0, 2, 4}, Speed: []int{4, 2, 1}},
			1,
		},
		{

			I{Target: 20, Position: []int{6, 2, 17}, Speed: []int{3, 9, 2}},
			2,
		},
		{
			I{Target: 10, Position: []int{8, 3, 7, 4, 6, 5}, Speed: []int{4, 4, 4, 4, 4, 4}},
			6,
		},
	}
	for _, tc := range testCases {
		result := CarFleet(tc.Input.Target, tc.Input.Position, tc.Input.Speed)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
