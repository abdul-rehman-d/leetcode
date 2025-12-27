package min_stack

import (
	"testing"
)

type TestCase struct {
	Input  string
	Output int
}

// Test case
// ["MinStack","push","push","push","getMin","pop","top","getMin"] -> functions
// [[],[-2],[0],[-3],[],[],[],[]] -> args
// [null,null,null,null,-3,null,0,-2] -> output

// ["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
// [[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
// [null,null,null,null,null,0,null,0,null,0,null,2]

func TestMain(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	result := obj.GetMin()
	if result != -3 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", -3, result)
		t.Fail()
	}

	obj.Pop()
	result = obj.Top()
	if result != 0 {
		t.Logf("\nTop failed: \nExpected: %+v\nActual: %+v", 0, result)
		t.Fail()
	}

	result = obj.GetMin()
	if result != -2 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", -2, result)
		t.Fail()
	}

	obj2 := Constructor()
	obj2.Push(2)
	obj2.Push(0)
	obj2.Push(3)
	obj2.Push(0)
	result = obj2.GetMin()
	if result != 0 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", 0, result)
		t.Fail()
	}
	obj2.Pop()
	result = obj2.GetMin()
	if result != 0 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", 0, result)
		t.Fail()
	}
	obj2.Pop()
	result = obj2.GetMin()
	if result != 0 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", 0, result)
		t.Fail()
	}
	obj2.Pop()
	result = obj2.GetMin()
	if result != 2 {
		t.Logf("\nGetMin failed: \nExpected: %+v\nActual: %+v", 2, result)
		t.Fail()
	}
}
