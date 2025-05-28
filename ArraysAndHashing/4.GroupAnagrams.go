package ArraysAndHashing

import (
	"reflect"
	"strings"
)

func groupAnagrams(str []string) [][]string {

	var m []map[string]int
	var ansIndex [][]int // stores the same index
	for _, s := range str {
		sepStr := strings.Split(s, "")
		eachMap := make(map[string]int)

		for _, c := range sepStr {
			if eachMap[c] == 0 {
				eachMap[c] = 1
			} else {
				eachMap[c] = eachMap[c] + 1
			}
		}
		m = append(m, eachMap)
	}

	if len(str) == 1 {
		return [][]string{{str[0]}}
	}

	for i, eachWord := range m {
		var group []int
		if eachWord != nil {
			group = append(group, i)
		}

		for j := i + 1; j < len(m); j++ {
			if reflect.DeepEqual(eachWord, m[j]) && m[j] != nil {
				group = append(group, j)
				m[j] = nil
			}
		}
		ansIndex = append(ansIndex, group)
	}

	var ans [][]string
	for _, i := range ansIndex {
		if i != nil {
			var groupStr []string
			for _, each := range i {
				groupStr = append(groupStr, str[each])
			}

			ans = append(ans, groupStr)
		}
	}

	return ans
}
