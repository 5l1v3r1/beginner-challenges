package main

import "fmt"

func main() {
	p, _ := ParsePolynomial("x + 3")
	p1, _ := ParsePolynomial("x + 1")
	p2, _ := ParsePolynomial("x - 5")
	p3, _ := ParsePolynomial("x + 9")
	fmt.Println(p.Multiply(p1).Multiply(p2).Multiply(p3))
}
