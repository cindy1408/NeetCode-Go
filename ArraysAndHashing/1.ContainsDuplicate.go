package ArraysAndHashing

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, n := range nums {
		if m[n] == 0 {
			m[n] = 1
		} else {
			return true
		}
	}

	return false
}
