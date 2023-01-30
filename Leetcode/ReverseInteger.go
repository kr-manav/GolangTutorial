package main

import "fmt"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		z := x % 10
		x = x - z
		x = x / 10
		ans = ans*10 + z
	}
	if ans > 2147483648 || ans < -2147483648 {
		return 0
	} else {
		return ans
	}
}

func main() {
	var s int
	fmt.Scanln(&s)
	fmt.Println("Reverse Number is : ", reverse(s))
}
