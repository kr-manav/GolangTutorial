package main

import "fmt"

var value = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	number := value[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		if value[string(s[i+1])] > value[string(s[i])] {
			number = number - value[string(s[i])]
		} else {
			number = number + value[string(s[i])]
		}
	}
	return number
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("Roman to Integer Value is : ", romanToInt(input))
}
