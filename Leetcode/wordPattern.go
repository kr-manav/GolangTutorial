package main

import (
	"fmt"
	"strings"
)

func checkValue(s string, i string, m map[string]string) bool {
	for k, v := range m {
		if k == i {
			continue
		}
		if v == s {
			fmt.Println(s, v)
			return true
		}
	}
	return false
}

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	var m = make(map[string]string)
	for i := 0; i < len(pattern); i++ {

		if m[string(pattern[i])] == "" && !checkValue(string(strs[i]), string(pattern[i]), m) {
			m[string(pattern[i])] = string(strs[i])
		} else {
			if m[string(pattern[i])] != string(strs[i]) || checkValue(string(strs[i]), string(pattern[i]), m) {
				return false
			}
		}
	}

	return true
}

func main() {
	var pattern string
	var s string
	pattern = "aaa"
	s = "cat cat cat"
	fmt.Println("Output for pattern matching: ", wordPattern(pattern, s))
}
