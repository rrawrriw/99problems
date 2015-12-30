package nintynine

type (
	Tree struct {
		Value string
		Left  *Tree
		Right *Tree
	}
)

func newBalTree(nodes int) *Tree {
	var t *Tree
	for x := 0; x < nodes; x++ {
		t = insertX(t, x)
	}

	return t
}

func insertX(t *Tree, node int) *Tree {
	if t == nil {
		return &Tree{"X", nil, nil}
	}
	if (node % 2) == 0 {
		t.Left = insertX(t.Left, node)
		return t
	}

	t.Right = insertX(t.Right, node)
	return t
}

func isSymmetric(t *Tree) bool {
	return cmp(t.Left, t.Right)
}

func cmp(t1, t2 *Tree) bool {
	c1, c2 := walker(t1), walker(t2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if (v1 == nil && v2 != nil) ||
			(v1 != nil && v2 == nil) {
			return false
		}

	}

	return true

}

func walker(t *Tree) chan *Tree {
	c := make(chan *Tree)
	go func() {
		walk(t, c)
		close(c)
	}()
	return c
}

func walk(t *Tree, c chan *Tree) {
	if t == nil {
		return
	}

	walk(t.Left, c)
	c <- t.Left
	walk(t.Right, c)
}
