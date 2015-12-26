package nintynine

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	l := []string{"a", "b", "c", "d"}
	i := size(l)
	fmt.Println(i)

	l = []string{}
	i = size(l)
	fmt.Println(i)
}

func size(l []string) int {
	var i int
	for i, _ = range l {
	}

	if i > 0 {
		i = i + 1
	}
	return i
}
