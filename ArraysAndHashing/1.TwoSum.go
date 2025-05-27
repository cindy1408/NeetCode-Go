package ArraysAndHashing

import (
	"slices"
)

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		var diff = target - num
		removedNum := make([]int, len(nums))
		copy(removedNum, nums)
		removedNum = append(removedNum[:i], removedNum[i+1:]...)
		if slices.Contains(removedNum, diff) {
			for j, n := range nums {
				if n == diff && j != i {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

func twoSumSolution(nums []int, target int) []int {
	for j, n := range nums {
		diff := target - n
		if slices.Contains(nums, diff) {
			for i := j + 1; i < len(nums); i++ {
				if nums[i] == diff {
					return []int{j, i}
				}
			}
		}
	}
	return nil
}
