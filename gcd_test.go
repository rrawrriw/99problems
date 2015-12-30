package nintynine

import "testing"

func TestGcd1(t *testing.T) {
	if r := gcd(20536, 7826); 2 != r {
		t.Fatal("Expect 2 was", r)
	}
}

func TestGcd2(t *testing.T) {
	if r := gcd(13, 27); 1 != r {
		t.Fatal("Expect 1 was", r)
	}
}

func gcd(a, b int) int {

	for a != 0 {
		//fmt.Printf("A: %v, B: %v\n", a, b)
		if a < b {
			b = b - a
			continue
		}

		a = a - b
	}

	return b
}
