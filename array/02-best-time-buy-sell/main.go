package main

import "fmt"

func maxProfit(prices []int) int {
	p, prev, price := 0, 0, 0
	hold := true
	for k, v := range prices {
		if k == 0 || prev > v {
			// descending
			if hold {
				p += (prev - price)
				hold = false
			}
		} else {
			// ascending
			if !hold {
				price = prev
				hold = true
			}
		}
		prev = v
	}
	if hold {
		p += (prev - price)
	}

	return p
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	answer := maxProfit(prices)

	fmt.Println(answer)
}
