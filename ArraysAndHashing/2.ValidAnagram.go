package ArraysAndHashing

import (
	"reflect"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")

	for _, n := range s1 {
		if m1[n] == 0 {
			m1[n] = 1
		} else {
			m1[n] = m1[n] + 1
		}
	}

	for _, n := range s2 {
		if m2[n] == 0 {
			m2[n] = 1
		} else {
			m2[n] = m2[n] + 1
		}
	}

	return reflect.DeepEqual(m1, m2)
}
