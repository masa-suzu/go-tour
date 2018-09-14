package tour

import "strings"

// WordCount counts up words that are simply divided by white space.
func WordCount(s string) map[string]int {
	count := map[string]int{}

	for _, v := range strings.Split(s, " ") {
		if _, ok := count[v]; !ok {
			count[v] = 0
		}
		count[v]++
	}
	return count
}
