package nintynine

import (
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {
	l := []string{"a", "b", "b", "b", "a", "c", "c"}
	r := pack(l)
	fmt.Println(r)
}

func pack(l []string) [][]string {
	n := [][]string{}
	last := ""

	for _, e := range l {
		if last != e {
			n = append(n, []string{})
		}

		n[len(n)-1] = append(n[len(n)-1], e)
		last = e
	}

	return n
}
