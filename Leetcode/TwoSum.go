package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans := []int{i, j}
				return ans
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{5, 10, 15, 20, 50}
	target := 35
	fmt.Println(twoSum(nums, target))
}
