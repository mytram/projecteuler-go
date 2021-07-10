package problem

import (
	"fmt"
	"math"
	"sort"
)

func nToDigits(n uint64) (digits []int) {
	for ; n > 0; n /= 10 {
		digits = append(digits, int(n%10))
	}

	return digits
}

func digitsToN(digits []int) (n uint64) {
	for i := len(digits) - 1; i >= 0; i-- {
		n += uint64(math.Pow10(i)) * uint64(digits[len(digits)-1-i])
	}

	return n
}

func normalizedCubeKey(cube uint64) uint64 {
	digits := nToDigits(cube)

	// sort the digits in descendingly order so that zeros are at the end and
	// therefore they will not be lost in the converted number

	sort.Slice(digits, func(i, j int) bool { return digits[i] > digits[j] })

	return digitsToN(digits)
}

func findCubesOfPerms(limit int) []uint64 {
	cubes := make(map[uint64][]uint64)

	var k uint64

	for n := uint64(1); len(cubes[k]) != limit; n++ {
		cube := n * n * n
		k = normalizedCubeKey(cube)
		cubes[k] = append(cubes[k], cube)
	}

	return cubes[k]
}

// TODO:
// Prove: this has exactly 5 perms
func Solve062() (string, uint64) {
	cubes := findCubesOfPerms(5)
	fmt.Println(cubes)
	return "062", cubes[0]
}
