package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var n int
	var r strings.Builder
	s := strconv.Itoa(x)
	r.Grow(len(s))
	if x < 0 {
		n = 1
		r.Write([]byte{'-'})
	}
	for i := len(s) - 1; i >= n; i-- {
		r.Write([]byte{s[i]})
	}

	res, err := strconv.Atoi(r.String())
	if err != nil || res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func main() {
	x := 1534236469

	res := reverse(x)

	fmt.Println(res)
}
