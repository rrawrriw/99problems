package nintynine

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	l := []string{"a", "b", "b", "b", "a", "c", "c"}
	r := encode(l)
	fmt.Println(r)
}

func encode(l []string) []map[string]int {
	n := []map[string]int{}
	last := ""

	for _, e := range l {
		if last != e {
			n = append(n, map[string]int{})
		}

		n[len(n)-1][e] = n[len(n)-1][e] + 1
		last = e
	}

	return n
}
