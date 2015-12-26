package nintynine

import (
	"fmt"
	"testing"
)

func TestLast(t *testing.T) {
	l := []string{"a", "b", "c", "d"}
	i := last(l)
	fmt.Println(i)

	l = []string{}
	i = last(l)
	fmt.Println(i)
}

func last(l []string) string {
	if len(l) == 0 {
		return ""
	}
	return l[len(l)-1]
}
