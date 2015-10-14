package main

import "fmt"

func main() {
	l := []string{"a", "b", "c"}
	v := duplicate(l)
	fmt.Println(v)
}

func duplicate(l []string) []string {
	n := []string{}
	for _, e := range l {
		n = append(n, e, e)
	}

	return n
}
