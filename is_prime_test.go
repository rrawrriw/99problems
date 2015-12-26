package nintynine

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPrime(22801763489))
}

func isPrime(n uint64) bool {

	if n <= 1 {
		return false
	}

	max := uint64(n + 1)
	s := make([]bool, max)

	for i := range s {
		s[i] = true
	}

	uB := uint64(math.Sqrt(float64(max)))
	for i := uint64(2); i <= uB; i++ {
		if !s[i] {
			continue
		}

		for j := i; (i * j) < max; j++ {
			s[i*j] = false
		}

	}

	return s[n]
}
