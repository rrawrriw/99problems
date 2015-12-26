package nintynine

import (
	"fmt"
	"testing"
)

func TestDrop(t *testing.T) {
	l := []string{"a", "b", "c", "d", "e", "f"}
	v := drop(l, 2)
	fmt.Println(v)
}

func drop(li []string, i int) []string {
	if len(li)-1 < i {
		return []string{}
	}

	l := make([]string, len(li))
	copy(l, li)

	ind := []int{}
	lo := len(l) / i
	for x := 1; x <= lo; x++ {
		ind = append(ind, (x*i)-x)
	}

	for _, e := range ind {
		l = append(l[:e], l[e+1:]...)
	}

	return l
}
