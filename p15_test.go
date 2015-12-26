package nintynine

import (
	"fmt"
	"testing"
)

func TestReplicate(t *testing.T) {
	l := []string{"a", "b", "c"}
	v := replicate(l, 3)
	fmt.Println(v)
}

func replicate(l []string, c int) []string {
	n := []string{}
	for _, e := range l {
		for x := 0; x < c; x++ {
			n = append(n, e)
		}
	}

	return n
}
