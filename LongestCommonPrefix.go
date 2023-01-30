package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	checkStr := strs[0]

	for _, v := range strs {
		if len(v) < len(checkStr) {
			checkStr = v
		}
	}
	var current string
	var ans string
	for i := 0; i <= len(checkStr); i++ {
		ans = current
		current = checkStr[:i]
		for _, v := range strs {
			if current != v[:i] {
				return ans
			}
		}
	}
	return current
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Common Prefix is : ", longestCommonPrefix(strs))
}
