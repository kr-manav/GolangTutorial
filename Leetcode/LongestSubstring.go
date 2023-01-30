package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var mx int
	for i := 0; i < len(s); i++ {
		var m = make(map[byte]int)
		var j int
		for j = i; j < len(s); j++ {
			if m[s[j]] == 1 {
				break
			}
			m[s[j]]++
		}
		if mx < j-i {
			mx = j - i
		}
	}
	return mx
}

func main() {
	s := "dvdf"
	fmt.Println("Length of longest substring is : ", lengthOfLongestSubstring(s))
}
