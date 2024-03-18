package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func myAtoi(s string) int {
	var r strings.Builder
	m := map[rune]uint{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	minInt := math.MinInt32
	maxInt := math.MaxInt32

	s = strings.Trim(s, " ")
	neg := false
	dig := false
	for _, v := range s {
		isDig := slices.Contains(digits, v)
		if dig {
			if isDig {
				r.Write([]byte(string(v)))
			} else {
				break
			}
		} else if v == '+' {
			dig = true
			continue
		} else if v == '-' {
			neg = true
			dig = true
			continue
		} else if isDig {
			dig = true
			r.Write([]byte(string(v)))
		} else {
			break
		}
	}
	sConv := r.String()

	var res uint64
	for k, v := range sConv {
		res += uint64(math.Pow(float64(10), float64(len(sConv)-k-1))) * uint64(m[v])
		if res > uint64(maxInt) && !neg {
			return maxInt
		}
		if res > -uint64(minInt) && neg {
			return minInt
		}
	}
	if neg {
		res = -res
	}

	return int(res)
}

func main() {
	s := "-91283472332"

	res := myAtoi(s)

	fmt.Println(res)
}
