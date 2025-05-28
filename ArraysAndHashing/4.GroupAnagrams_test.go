package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestGroupAnagramsExample1(t *testing.T) {
	lists := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	ans := groupAnagrams(lists)

	expected := [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}

	sortedAns := sortSlice(ans)
	sortedExpected := sortSlice(expected)

	assert.Equal(t, sortedAns, sortedExpected)
}

func TestGroupAnagramsExample2(t *testing.T) {
	lists := []string{"x"}
	ans := groupAnagrams(lists)

	assert.Equal(t, [][]string{{"x"}}, ans)
}

func TestGroupAnagramsExample3(t *testing.T) {
	lists := []string{""}
	ans := groupAnagrams(lists)

	assert.Equal(t, [][]string{{""}}, ans)
}

func sortSlice(ans [][]string) [][]string {
	for _, str := range ans {
		sort.Strings(str)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return nil
}
