package ext

import (
	"strconv"
	"strings"
)

// Int64Join joins slice of int64
func Int64Join(a []int64, sep string) string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(s, sep)
}

// RmDup remove duplicate element of []string
func RmDup(list []string) []string {
	if len(list) == 0 {
		return []string{}
	}
	m := make(map[string]bool)
	result := make([]string, 0)
	for _, v := range list {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}
