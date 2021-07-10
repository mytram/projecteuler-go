# projecteuler-go

## Problem 058

This problem is to generate the Ulam spiral. The Ulam spiral is generate by placing natural numbers into squares anticlockwise. For example, below is the first three squares of the Ulam spiral.

```
37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49
```

### Data representation

For each number, a 2d position coordinate can be systematically assigned to each number in the spiral in the following way:

- the origin position is `(0, 0)`
- Assume that the spiral's head is at `(x, y)`. The next number can be imagined as moving in the following direction relative to the head:
  - right, the next number's position is `(x + 1, y)`
  - up, `(x, y + 1)`
  - left, `(x - 1, y)`
  - down, `(x, y - 1)`

For example, using this coordinate system, we can represent the first 9 numbers in the format `(n, x, y)`, where `n` is the number, as follows:

`(1, 0, 0), (2, 1, 0), (3, 1, 1), (4, 0, 1), (5, -1, 1), (6, -1, 0), (7, -1, -1), (8, 0, -1), (9, 1, -1)`

### Some properties of the spiral and the coordinate system that are useful for the solution

- The lengths of the squares are the odd numbers, e.g. `1, 3, 5, 6...`, and therefore `next length = current length + 2`
- Each square of the spiral ends at the square of the length of the square
- The coordinates on the two diagonals have the same absolute values. In particular, `x` and `y` are equal on the top right to bottom left diagonal, and `x + y` is `0` on the other
- With the above knowledge, the coordinate of the ending number of a square is `(x = (length - 1) / 2, y = -x)` using only the `current length`. For example, the number `9` is the ending number of the second square of length `3`. Its coordinate is `(x = (3 - 1)/2 = 1, y = -x = -1)`
- From the last number of the previous square, the head moves in the following ways to generate the next square
  - 1 step to the right
  - the `next length - 2` steps up
  - the `next length - 1` steps to the left
  - the `next length - 1` steps down
  - the `next length - 1` steps to the right again

### The algorithm

- Generate a spiral by complete squares
- Use the coordinate to check if the number is on the diagonal and remember the count
- If the number is on the diagonal, check if it is a prime using `math.Big` and remember the count
- For each spiral square, stop the program if the ratio of prime numbers on diagonals have dropped below `10%`
- The current length is the solution

## Problem 059

The plain text is

```
An extract taken from the introduction of one of Euler's most celebrated papers, "De summis serierum reciprocarum" [On the sums of series of reciprocals]: I have recently found, quite unexpectedly, an elegant expression for the entire sum of this series 1 + 1/4 + 1/9 + 1/16 + etc., which depends on the quadrature of the circle, so that if the true sum of this series is obtained, from it at once the quadrature of the circle follows. Namely, I have found that the sum of this series is a sixth part of the square of the perimeter of the circle whose diameter is 1; or by putting the sum of this series equal to s, it has the ratio sqrt(6) multiplied by s to 1 of the perimeter to the diameter. I will soon show that the sum of this series to be approximately 1.644934066842264364; and from multiplying this number by six, and then taking the square root, the number 3.141592653589793238 is indeed produced, which expresses the perimeter of a circle whose diameter is 1. Following again the same steps by which I had arrived at this sum, I have discovered that the sum of the series 1 + 1/16 + 1/81 + 1/256 + 1/625 + etc. also depends on the quadrature of the circle. Namely, the sum of this multiplied by 90 gives the biquadrate (fourth power) of the circumference of the perimeter of a circle whose diameter is 1. And by similar reasoning I have likewise been able to determine the sums of the subsequent series in which the exponents are even numbers.
```

## Problem 062
