package nintynine

import (
	"fmt"
	"testing"
)

func TestIsPal(t *testing.T) {
	l := []string{"a", "b", "a"}
	r := isPal(l)
	fmt.Println(r)

	l = []string{"a", "b", "c"}
	r = isPal(l)
	fmt.Println(r)

}

func isPal(l []string) bool {
	return equal(l, rev(l))
}

func equal(l1, l2 []string) bool {
	if len(l1) != len(l2) {
		return false
	}

	for x := 0; x < len(l1); x++ {
		if l1[x] != l2[x] {
			return false
		}
	}

	return true
}
