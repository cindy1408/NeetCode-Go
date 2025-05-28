package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumsExample1(t *testing.T) {
	nums := []int{3, 4, 5, 6}
	target := 7
	ans := twoSums(nums, target)

	assert.Equal(t, []int{0, 1}, ans)
}

func TestTwoSumsExample2(t *testing.T) {
	nums := []int{4, 5, 6}
	target := 10
	ans := twoSums(nums, target)

	assert.Equal(t, []int{0, 2}, ans)
}

func TestTwoSumsExample3(t *testing.T) {
	nums := []int{5, 5}
	target := 10
	ans := twoSums(nums, target)

	assert.Equal(t, []int{0, 1}, ans)
}
