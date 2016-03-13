package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// A Polynomial represents a polynomial on a single variable of the form a+bx+cx^2+..., where each
// of the entries in the underlying []float64 corresponds to a coefficient in the polynomial.
// For instance, p[4] is the coefficient for the x^4 term in a Polynomial p.
type Polynomial []float64

// ParsePolynomial returns the polynomial for a given human-readable polynomial string.
func ParsePolynomial(s string) (Polynomial, error) {
	parts := strings.Split(s, " ")
	if len(parts)%2 == 0 {
		return nil, errors.New("invalid polynomial: " + s)
	}
	res := Polynomial{}
	for i := 0; i < len(parts); i += 2 {
		var term Polynomial
		var err error
		if i == 0 {
			term, err = parseTerm("+", parts[i])
		} else {
			term, err = parseTerm(parts[i-1], parts[i])
		}
		if err != nil {
			return nil, err
		}
		res = res.Add(term)
	}
	return res, nil
}

// Add returns the sum of this polynomial with another one.
func (p Polynomial) Add(p1 Polynomial) Polynomial {
	maxLen := len(p)
	if len(p1) > len(p) {
		maxLen = len(p1)
	}
	res := make(Polynomial, maxLen)
	for i := range res {
		var pCoeff, p1Coeff float64
		if i < len(p) {
			pCoeff = p[i]
		}
		if i < len(p1) {
			p1Coeff = p1[i]
		}
		res[i] = pCoeff + p1Coeff
	}
	return res
}

// Scale returns the product of this polynomial with a constant.
func (p Polynomial) Scale(f float64) Polynomial {
	res := make(Polynomial, len(p))
	copy(res, p)
	for i, x := range res {
		res[i] = x * f
	}
	return res
}

// RaisePower is equivalent to multiplying this polynomial by x^n.
func (p Polynomial) RaisePower(n int) Polynomial {
	res := make(Polynomial, len(p)+n)
	copy(res[n:], p)
	return res
}

// Multiply multiplies this polynomial by another polynomial.
func (p Polynomial) Multiply(p1 Polynomial) Polynomial {
	res := Polynomial{}
	for i, coeff := range p {
		res = res.Add(p1.RaisePower(i).Scale(coeff))
	}
	return res
}

// String returns a human-readable version of this polynomial.
func (p Polynomial) String() string {
	terms := make([]string, len(p))
	for i, coefficient := range p {
		coefficientStr := strconv.FormatFloat(coefficient, 'f', -1, 64)
		terms[len(p)-i-1] = coefficientStr + "x^" + strconv.Itoa(i)
	}
	return strings.Join(terms, " + ")
}

func parseTerm(operation, term string) (Polynomial, error) {
	var sign float64
	switch operation {
	case "+":
		sign = 1
	case "-":
		sign = -1
	default:
		return nil, errors.New("unknown operation: " + operation)
	}

	expr := regexp.MustCompile("^(-?)([0-9\\.]*)(x(\\^([0-9]*))?)?$")
	match := expr.FindStringSubmatch(term)
	if match == nil {
		return nil, errors.New("invalid term: " + term)
	}

	if match[1] == "-" {
		sign *= -1
	}

	coefficient, err := strconv.ParseFloat(match[2], 64)
	if match[2] == "" {
		coefficient = 1
		err = nil
	}
	if err != nil {
		return nil, err
	}

	exponent, err := strconv.Atoi(match[5])
	if match[3] == "" {
		exponent = 0
		err = nil
	} else if match[4] == "" {
		exponent = 1
		err = nil
	}
	if err != nil {
		return nil, err
	}

	p := make(Polynomial, exponent+1)
	p[exponent] = coefficient * sign
	return p, nil
}
