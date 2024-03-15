package main

import (
	"fmt"
)

func firstUniqueChar(s string) int {
	m := make(map[byte]struct{}, 26)
	p := make(map[byte]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := m[s[i]]; ok {
			delete(p, s[i])
		} else {
			m[s[i]] = struct{}{}
			p[s[i]] = i
		}
	}
	r := len(s)
	for _, v := range p {
		r = min(r, v)
	}

	if r == len(s) {
		return -1
	}

	return r
}

func main() {
	s := "leetcode"

	res := firstUniqueChar(s)

	fmt.Println(res)
}
