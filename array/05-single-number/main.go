package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0

	// res, cnext, cprev := 0, false, true
	// slices.Sort(nums)
	//
	//	for i := 0; i < len(nums); i++ {
	//		cnext = (i == len(nums)-1 || nums[i] != nums[i+1])
	//		if cprev && cnext {
	//			res = nums[i]
	//			break
	//		}
	//		cprev = cnext
	//	}
	//
	// return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)

	fmt.Println(res)
}
