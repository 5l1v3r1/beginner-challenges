package main

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const rootSearchIterations = 64
const edgeSearchIterations = 512

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

// Evaluate plugs a value into this polynomial.
func (p Polynomial) Evaluate(x float64) float64 {
	var res float64
	xTerm := 1.0
	for _, coefficient := range p {
		res += coefficient * xTerm
		xTerm *= x
	}
	return res
}

// Degree returns the degree of the polynomial, or -1 if this is the zero polynomial.
func (p Polynomial) Degree() int {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != 0 {
			return i
		}
	}
	return -1
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

// Derivative returns the derivative of this polynomial.
func (p Polynomial) Derivative() Polynomial {
	if len(p) == 0 {
		return p
	}
	res := make(Polynomial, len(p)-1)
	for i := range res {
		res[i] = p[i+1] * float64(i+1)
	}
	return res
}

// Roots returns the roots of the polynomial.
func (p Polynomial) Roots() []float64 {
	p = p[:p.Degree()+1]

	if len(p) == 0 || len(p) == 1 {
		return nil
	} else if len(p) == 2 {
		if p[1] == 0 {
			p[:1].Roots()
		}
		return []float64{-p[0] / p[1]}
	}

	res := []float64{}

	criticalPoints := p.Derivative().Roots()
	sort.Float64s(criticalPoints)
	if len(criticalPoints) == 0 {
		if p.Evaluate(0) == 0 {
			return []float64{0}
		}
		res = append(res, p.rootAfterFirstCriticalPoint(0)...)
		res = append(res, p.rootBeforeFirstCriticalPoint(0)...)
		return res
	}

	for i := 0; i < len(criticalPoints)-1; i++ {
		p1, p2 := criticalPoints[i], criticalPoints[i+1]
		res = append(res, p.rootBetweenCriticalPoints(p1, p2)...)
	}
	res = append(res, p.rootBeforeFirstCriticalPoint(criticalPoints[0])...)
	res = append(res, p.rootAfterFirstCriticalPoint(criticalPoints[len(criticalPoints)-1])...)
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

func (p Polynomial) rootBetweenCriticalPoints(p1, p2 float64) []float64 {
	value1 := p.Evaluate(p1)
	value2 := p.Evaluate(p2)
	if value1 < 0 && value2 < 0 {
		return []float64{}
	} else if value1 > 0 && value2 > 0 {
		return []float64{}
	} else if value1 > 0 {
		return p.rootBetweenCriticalPoints(p2, p1)
	}

	for i := 0; i < rootSearchIterations; i++ {
		mid := (p1 + p2) / 2
		value := p.Evaluate(mid)
		if value == 0 {
			return []float64{mid}
		} else if value < 0 {
			p1 = mid
		} else {
			p2 = mid
		}
	}
	return []float64{(p1 + p2) / 2}
}

func (p Polynomial) rootBeforeFirstCriticalPoint(x float64) []float64 {
	beforeDerivative := p.Derivative().Evaluate(x - 1.0)
	criticalValue := p.Evaluate(x)
	if (beforeDerivative < 0 && criticalValue > 0) ||
		(beforeDerivative > 0 && criticalValue < 0) {
		return []float64{}
	}
	difference := 1.0
	for i := 0; i < edgeSearchIterations; i++ {
		value := p.Evaluate(x - difference)
		if value < 0 && criticalValue > 0 || value > 0 && criticalValue < 0 {
			break
		}
		difference *= 2
	}
	return p.rootBetweenCriticalPoints(x-difference, x)
}

func (p Polynomial) rootAfterFirstCriticalPoint(x float64) []float64 {
	newP := p.Scale(1)
	for i, x := range newP {
		if i%2 != 0 {
			newP[i] = -x
		}
	}
	root := newP.rootBeforeFirstCriticalPoint(-x)
	for i, x := range root {
		root[i] = -x
	}
	return root
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
