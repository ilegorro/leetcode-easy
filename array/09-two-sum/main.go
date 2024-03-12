package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// the fastest solution
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for k, v := range nums {
		if km, ok := m[target-v]; ok && k != km {
			return []int{k, km}
		}
	}

	// slower but more memory efficient
	// for k, v := range nums {
	// 	if res := slices.Index(nums[k+1:], target-v); res != -1 {
	// 		return []int{k, res + k + 1}
	// 	}
	// }

	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	return nil
}

func main() {
	nums := []int{1, 3, 4, 2}
	target := 6
	res := twoSum(nums, target)

	fmt.Println(res)
}
