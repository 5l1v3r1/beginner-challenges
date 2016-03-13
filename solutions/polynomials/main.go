package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	parts := []string{"x - 3", "x - 2", "x - 1", "x", "x + 1", "x + 2", "x + 3"}
	res := Polynomial{1}
	for _, part := range parts {
		p, _ := ParsePolynomial(part)
		res = res.Multiply(p)
	}
	factorMe := res.Add(Polynomial{1})
	fmt.Println("Roots of", factorMe, "are", factorMe.Roots())

	// Figure out the % of roots of 8th degree polynomials that are real.
	rand.Seed(time.Now().UnixNano())
	totalRootCount := 0
	totalPossibleRootCount := 0
	for i := 0; i < 100000; i++ {
		randPoly := make(Polynomial, 9)
		for i := range randPoly {
			randPoly[i] = rand.Float64()*2 - 1
		}
		totalRootCount += len(randPoly.Roots())
		totalPossibleRootCount += len(randPoly) - 1
	}
	fmt.Println("Polynomials had", 100*float64(totalRootCount)/float64(totalPossibleRootCount),
		"percent real roots.")
}
