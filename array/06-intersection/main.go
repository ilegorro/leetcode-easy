package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}
	for k1, v1 := range m1 {
		v2 := m2[k1]
		for i := 0; i < min(v1, v2); i++ {
			res = append(res, k1)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	res := intersect(nums1, nums2)

	fmt.Println(res)
}
