package main

import "fmt"

func isValid(s string) bool {
	var checkStr string
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			checkStr = checkStr + string(s[i])
		} else {
			if len(checkStr) > 0 {
				if s[i] == ')' {
					if checkStr[len(checkStr)-1] == '(' {
						checkStr = checkStr[:len(checkStr)-1]
					} else {
						return false
					}
				} else if s[i] == ']' {
					if checkStr[len(checkStr)-1] == '[' {
						checkStr = checkStr[:len(checkStr)-1]
					} else {
						return false
					}
				} else if s[i] == '}' {
					if checkStr[len(checkStr)-1] == '{' {
						checkStr = checkStr[:len(checkStr)-1]
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	if len(checkStr) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println("Parthesis Valid : ", isValid(s))
}
