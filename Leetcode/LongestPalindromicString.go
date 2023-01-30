package main

import "fmt"

func longestPalindrome(s string) string {
	var ans string
	for i := 0; i <= len(s); i++ {
		j := i
		for j = i; j <= len(s); j++ {
			checkStr := s[i:j]
			flag := true
			for k := 0; k < len(checkStr)/2; k++ {
				if checkStr[k] != checkStr[len(checkStr)-k-1] {
					flag = false
					break
				}
			}
			if len(ans) < len(checkStr) && flag == true {
				ans = checkStr
			}
		}
	}
	return ans
}

func main() {
	s := "babad"
	fmt.Println("Longest Palindromic String is : ", longestPalindrome(s))
}
