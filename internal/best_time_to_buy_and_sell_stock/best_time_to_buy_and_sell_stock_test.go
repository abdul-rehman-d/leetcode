package best_time_to_buy_and_sell_stock_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/best_time_to_buy_and_sell_stock"
)

func TestMain(t *testing.T) {
    inputs := [4][]int{
        {7,1,5,3,6,4},
        {7,6,4,3,1},
        {1,2},
        {1,2,4,2,5,7,2,4,9,0,9},
    }
    outputs := [4]int{
        5,
        0,
        1,
        9,
    }

    for idx, input := range inputs {
		result := best_time_to_buy_and_sell_stock.MaxProfit(input)
		if (result != outputs[idx]) {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
    }
}


