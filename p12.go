package main

import "fmt"

type (
	one struct {
		c string
	}
	many struct {
		c string
		n int
	}
)

func main() {
	l := []interface{}{
		many{"a", 4},
		one{"b"},
		many{"c", 2},
		many{"a", 2},
		one{"d"},
	}
	r := decode(l)
	fmt.Println(r)
}

func decode(l []interface{}) []string {
	n := []string{}
	
	for _, e := range l {
		switch v := e.(type) {
		case one:
			n = append(n, v.c)
		case many:
			for x := 0; x < v.n; x++ {
				n = append(n, v.c)
			}
		}
	
	}
	
	return n
}

