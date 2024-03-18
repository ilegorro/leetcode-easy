package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var r strings.Builder
	maxI := 200
	for _, v := range strs {
		maxI = min(maxI, len(v))
	}
	for i := 0; i < maxI; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return r.String()
			}
		}
		r.Write([]byte{c})
	}

	return r.String()
}

func main() {
	strs := []string{"dog", "racecar", "car"}

	res := longestCommonPrefix(strs)

	fmt.Println(res)
}
