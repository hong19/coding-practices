package leetcode322

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	leastAmount := make([]int, amount+1)

	for i := range leastAmount {
		leastAmount[i] = -1
	}
	leastAmount[0] = 0

	for _, value := range coins {
		if value <= amount {
			leastAmount[value] = 1
		}
	}

	for current := 0; current < len(leastAmount); current++ {
		if leastAmount[current] == -1 {
			for _, coin := range coins {
				if (current - coin) <= 0 {
					continue
				}
				if leastAmount[current-coin] == -1 {
					continue
				}

				if leastAmount[current] == -1 {
					leastAmount[current] = leastAmount[coin] + leastAmount[current-coin]
				} else if (leastAmount[coin] + leastAmount[current-coin]) < leastAmount[current] {
					leastAmount[current] = leastAmount[coin] + leastAmount[current-coin]
				}
			}
		}
	}

	fmt.Printf("coins: %v, leastAmount: %v\n", coins, leastAmount)

	return leastAmount[amount]
}
