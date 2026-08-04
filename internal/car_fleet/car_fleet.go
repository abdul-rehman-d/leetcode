package carfleet

import (
	"slices"
)

type Car struct {
	P int
	S int
}

func CarFleet(target int, position []int, speed []int) int {
	l := len(position)
	if len(speed) != l {
		panic("wrong input")
	}

	cars := make([]Car, l)
	for i := range l {
		cars[i] = Car{
			S: speed[i],
			P: position[i],
		}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return b.P - a.P
	})

	stack := []float64{}
	// e.g. [(0,1), (3,3), (5,1), (8,4), (10,2)]
	// i | stack before comparison | stack after comparison
	// - | ----------------------- | ----------------------
	// 0 | []                      | [1]
	// 1 | [1, 1]                  | [1]
	// 2 | [1, 7]                  | [1, 7]
	// 3 | [1, 7, 3]               | [1, 7]
	// 4 | [1, 7, 12]              | [1, 7, 12]

	for _, car := range cars {
		t := float64(float64(target-car.P) / float64(car.S))
		stack = append(stack, t)

		if len(stack) >= 2 {
			last := stack[len(stack)-1]
			sLast := stack[len(stack)-2]
			if last <= sLast {
				// pop
				stack = stack[:len(stack)-1]
			}
		}

	}

	return len(stack)
}
