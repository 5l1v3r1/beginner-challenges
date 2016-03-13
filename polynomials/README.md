# Overview

This challenge is to create a suite of tools for manipulating polynomials. It is suggested that you make a `Polynomial` class for this project which implements a bunch of the below functionality.

# Features

Your suite of polynomial functions (e.g. a `Polynomial` class) should support the following functionality:

 * Parse a polynomial that a user types.
 * Output a polynomial as a human-readable string (implement the `toString()` method in Java).
 * Find or approximate the roots of a polynomial.
 * Add polynomials to other polynomials.
 * Multiply polynomials by a number (using the distributive law).
 * Multiply polynomials by other polynomials, yielding another polynomial (using the distributive law).
 * Graph a polynomial using some kind of GUI library.

## Polynomial string formats

Polynomials can be represented as strings of the form "a + bx + cx^2 + cx^3 ...". Coefficients may be omitted and minus signs may be used (e.g. "1 - x + x^2 - 3x^3"). If you want a challenge, make your parser able to handle weirder polynomial formats (e.g. out of order terms, multiple terms of the same degree, no spaces between terms, etc.).

## Polynomial roots

There are many ways to find the roots of polynomials. You might simply implement a brute force solution; such a thing is relatively easy for a computer. You might also try a binary search or some other method. You can probably come up with something clever, especially if you know calculus.

# Specific challenges

Once you've made your polynomial tools, try to find the real roots of the following polynomials:

```
2x^2 + 2x - 1
```

```
x^5 + 7x^4 - 179x^3 - 783x^2 + 1890x + 5400
```

```
(x-3)(x-2)(x-1)x(x+1)(x+2)(x+3) + 1
```

```
(10x^3 + 2x^2 - 1)(10x^2 + 2x) - 5(7x^3 + 3x^2 + 1x + 10)
```

Graph the following polynomials:

```
-x + 0.1666666667x^3 - 0.008333333333x^5 + 0.0001984126984x^7 - 0.000002777777778x^9
```

```
(x-3)(x-2)(x-1)x(x+1)(x+2)(x+3)
```
