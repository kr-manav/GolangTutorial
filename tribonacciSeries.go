package main

import "fmt"

func tribonacci(n int) int {
	x, y, z := 0, 1, 1
	ans := 0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		for i := 0; i < n-2; i++ {
			ans = x + y + z
			x = y
			y = z
			z = ans
		}
		return ans
	}
}

func main() {
	var input int
	fmt.Scanln(&input)
	fmt.Printf("Value of %d number in tribonacci series is : %d\n", input, tribonacci(input))
}
