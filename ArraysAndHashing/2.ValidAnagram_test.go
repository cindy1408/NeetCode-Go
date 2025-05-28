package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagramExample1(t *testing.T) {
	s1 := "racecar"
	s2 := "carrace"
	ans := isAnagram(s1, s2)

	assert.Equal(t, true, ans)
}

func TestIsAnagramExample2(t *testing.T) {
	s1 := "jar"
	s2 := "carrace"
	ans := isAnagram(s1, s2)

	assert.Equal(t, false, ans)
}
