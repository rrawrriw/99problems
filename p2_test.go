package nintynine

import (
	"fmt"
	"testing"
)

func TestLastTwo(t *testing.T) {
	l := []string{"a", "b", "c", "d"}
	i := lastTwo(l)
	fmt.Println(i)

	l = []string{"a"}
	i = lastTwo(l)
	fmt.Println(i)
}

func lastTwo(l []string) []string {
	if len(l) < 2 {
		return []string{}
	}
	return l[len(l)-2 : len(l)-1]
}
