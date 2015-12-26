package nintynine

import (
	"fmt"
	"testing"
)

func TestRev(t *testing.T) {
	l := []string{"a", "b", "c"}
	r := rev(l)
	fmt.Println(r)

}

func rev(l []string) []string {
	n := []string{}
	for x := len(l) - 1; x >= 0; x-- {
		n = append(n, l[x])
	}

	return n
}
