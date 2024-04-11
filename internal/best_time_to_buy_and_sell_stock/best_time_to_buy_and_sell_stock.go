package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
    start, end := 0, 1
    profit := 0

    for start < end && end < len(prices) {
        curr := prices[end] - prices[start]

        if curr > profit {
            profit = curr
        }

        if prices[start] < prices[end] {
            end++
        } else {
            start=end
            end++
        }
    }

    return profit
}

