package main

import "fmt"

func getArray(x int) []int {
	arr := []int{}
	for x != 0 {
		y := x % 10
		arr = append(arr, y)
		x = (x - y) / 10
	}
	return arr
}

func isPalindrome(x int) bool {
	var b []int
	//a := getForwardOrder(x, a)
	if x < 0 {
		return false
	}
	b = getArray(x)
	for i := 0; i < len(b); i++ {
		if b[i] != b[len(b)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	x := -121
	fmt.Println(isPalindrome(x))
}
