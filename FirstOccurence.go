package main

import "fmt"

func strStr(haystack string, needle string) int {
	for j := 0; j <= len(haystack)-len(needle); j++ {
		if needle == haystack[j:j+len(needle)] {
			return j
		}
	}
	return -1
}

func main() {
	haystack := "dogcatdog"
	needle := "og"
	fmt.Println("First Occurence is at Index : ", strStr(haystack, needle))
}
