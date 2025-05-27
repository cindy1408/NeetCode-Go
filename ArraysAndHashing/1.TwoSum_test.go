package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumExample1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSumSolution(nums, target)

	assert.Equal(t, []int{0, 1}, ans)
}

func TestTwoSumExample2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	ans := twoSumSolution(nums, target)

	assert.Equal(t, []int{1, 2}, ans)
}

func TestTwoSumExample3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	ans := twoSumSolution(nums, target)

	assert.Equal(t, []int{0, 1}, ans)
}
