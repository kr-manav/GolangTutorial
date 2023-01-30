package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	count := 0
	sum := 0
	neg := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		neg = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for sum <= dividend {
		count++
		sum = sum + divisor
	}
	count--
	if neg == false {
		if count > 2147483647 {
			return 2147483647
		} else if count < -2147483647 {
			return -2147483647
		} else {
			return count
		}

	} else {
		count = -count
		if count <= -2147483649 {
			return -2147483648
		} else {
			return count
		}
	}

}

func main() {
	var dividend, divisor int
	fmt.Scanln(&dividend)
	fmt.Scanln(&divisor)
	fmt.Println("Quotient is : ", divide(dividend, divisor))
}
