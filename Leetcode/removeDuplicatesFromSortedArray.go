package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	count := 0
	for i := 0; i < length; i++ {
		elem := nums[i]
		if elem != 500 {
			for j := i + 1; j < length; j++ {
				if nums[j] == elem && nums[j] != 500 {
					nums[j] = 500
					count++
				}
			}
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] == 500 {
			for j := i + 1; j < length; j++ {
				if nums[j] != 500 {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					break
				}
			}
		}
	}

	// fmt.Println(nums)
	return length - count
}

func main() {
	s := []int{0, 0, 1, 1, 1, 2, 2}
	fmt.Println("New Array is : ", s[:removeDuplicates(s)])
}
