package main

import "fmt"

func main() {
	parts := []string{"x - 3", "x - 2", "x - 1", "x", "x + 1", "x + 2", "x + 3"}
	res := Polynomial{1}
	for _, part := range parts {
		p, _ := ParsePolynomial(part)
		res = res.Multiply(p)
	}
	factorMe := res.Add(Polynomial{1})
	fmt.Println("Factors of", factorMe, "are", factorMe.Roots())
}
