package nintynine

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tree := newBalTree(4)
	fmt.Println(isSymmetric(tree))

	tree = newBalTree(5)
	fmt.Println(isSymmetric(tree))
}
