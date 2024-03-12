package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}) // an empty struct occupies zero bytes
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := containsDuplicate(nums)

	fmt.Println(res)
}
