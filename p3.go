package main

import "fmt"

func main() {
	l := []string{"a", "b", "c", "d"}
	i := at(l, 3)
	fmt.Println(i)

	l = []string{"a"}
	i = at(l, 2)
	fmt.Println(i)
}

func at(l []string, i int) string {
	if len(l) - 1 < i {
		return ""
	}
	return l[i]
}

