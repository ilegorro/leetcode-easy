package main

import (
	"fmt"
)

func moveZeros(nums []int) {
	// this solution is faster
	res := make([]int, 0)
	for _, v := range nums {
		if v != 0 {
			res = append(res, v)
		}
	}
	for i := len(res); i < len(nums); i++ {
		res = append(res, 0)
	}

	copy(nums, res)

	// this solution is more memory effective
	// l := len(nums)
	// res := nums
	// for i := 0; i < len(res); i++ {
	// 	if res[i] == 0 {
	// 		res = append(res[:i], res[i+1:]...)
	// 		i--
	// 	}
	// }
	// for i := len(res); i < l; i++ {
	// 	res = append(res, 0)
	// }
	// copy(nums, res)
}

func main() {
	nums := []int{0, 1, 0, 0, 3, 12}
	moveZeros(nums)

	fmt.Println(nums)
}
