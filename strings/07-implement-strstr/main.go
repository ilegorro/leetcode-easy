package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		r := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				r = false
				break
			}
		}
		if r {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "saadbutsad"
	needle := "sad"

	res := strStr(haystack, needle)

	fmt.Println(res)
}
