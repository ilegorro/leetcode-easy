package main

import "fmt"

func removeDuplicates(nums []int) int {
	t, prev := 0, 0
	for k, v := range nums {
		if v != prev || k == 0 {
			nums[t] = v
			prev = v
			t++
		}
	}

	return t
}

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4}
	answer := removeDuplicates(nums)

	fmt.Println(answer)
}
