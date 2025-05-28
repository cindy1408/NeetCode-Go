package ArraysAndHashing

func twoSums(nums []int, target int) []int {
	for i, n := range nums {
		diff := target - n
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == diff && j != i {
				return []int{i, j}
			}
		}
	}

	return nil
}
