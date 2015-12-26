package nintynine

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	l := []string{"a", "b", "b", "b", "a", "c", "c"}
	r := compress(l)
	fmt.Println(r)
}

func compress(l []string) []string {
	n := []string{}
	last := ""

	for _, i := range l {
		if last != i {
			n = append(n, i)
		}

		last = i
	}

	return n
}
