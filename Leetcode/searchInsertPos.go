package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums[0] > target || nums[0] == target {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	} else {
		for i := 1; i <= len(nums)-1; i++ {
			if nums[i] == target {
				return i
			} else {
				if nums[i-1] < target && nums[i] > target {
					return i
				}
			}
		}
	}

	return 0
}

func main() {
	nums := []int{1, 2, 4, 7, 8, 9}
	target := 6
	fmt.Println("Correct Position to Insert is at Index :", searchInsert(nums, target))
}
