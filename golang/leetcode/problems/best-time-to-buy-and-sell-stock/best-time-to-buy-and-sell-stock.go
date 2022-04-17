package best_time_to_buy_and_sell_stock

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	profit := 0
	lastPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if profit < prices[i]-lastPrice {
			profit = prices[i] - lastPrice
		}

		if lastPrice > prices[i] {
			lastPrice = prices[i]
		}
	}

	return profit
}
