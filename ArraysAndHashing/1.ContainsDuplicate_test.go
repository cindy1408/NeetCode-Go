package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicateExample1(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	ans := containsDuplicate(nums)

	assert.Equal(t, true, ans)
}

func TestContainsDuplicateExample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	ans := containsDuplicate(nums)

	assert.Equal(t, false, ans)
}
