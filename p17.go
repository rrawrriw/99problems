package main

import "fmt"

func main() {
	l := []string{"a", "b", "c"}
	v := drop(l, 1)
	fmt.Println(v)
}
