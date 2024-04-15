package strings

import "strings"

func StartsWithAny(array []string, s string) bool {
	for _, prefix := range array {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}
