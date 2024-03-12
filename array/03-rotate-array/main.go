package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	tail := nums[len(nums)-k:]
	head := nums[:len(nums)-k]
	res := append(tail, head...)
	copy(nums, res)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	rotate(nums, k)

	fmt.Println(nums)
}
