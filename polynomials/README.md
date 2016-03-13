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
