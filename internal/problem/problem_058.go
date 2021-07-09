package problem

import (
	"math/big"

	"github.com/mytram/projecteuler-go/internal/ulamspiral"
)

type counters struct {
	diagonalCount int
	primeCount    int
}

func (c *counters) count(n, x, y int) {
	if x != y && x+y != 0 {
		return
	}

	c.diagonalCount += 1

	if big.NewInt(int64(n)).ProbablyPrime(0) {
		c.primeCount += 1
	}
}

func (c *counters) done() bool {
	return c.primeCount > 0 && float32(c.primeCount)/float32(c.diagonalCount) < 0.1
}

func Solve058() (string, int) {
	c := &counters{1, 0}

	proc := func(n, x, y int) { c.count(n, x, y) }

	spiral := ulamspiral.New()

	for !c.done() {
		spiral.Circle(proc)
	}

	return "058", spiral.Length
}
