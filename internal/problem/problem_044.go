package problem

import (
	"fmt"

	"github.com/mytram/projecteuler-go/internal/problem/pentagonnumbers"
)

func Solve044() (string, uint) {
	minDiff := ^uint(0)
	stop := false

	pn := pentagonnumbers.New()

	for n := uint(1); !stop; n++ {
		nth := pn.Get(n)

		for i := n - 1; i > 0; i-- {
			ith := pn.Get(i)
			diff := nth - ith

			if diff >= minDiff {
				break
			}

			if !pn.IsPentagonNumber(diff) {
				continue
			}

			sum := nth + ith

			if !pn.IsPentagonNumber(sum) {
				continue
			}

			fmt.Printf("n=%d nth=%d i=%d nth=%d diff=%d diffN=%d sum=%d sumN=%d\n",
				n, nth, i,
				pn.Get(i),
				diff,
				pn.GetIndex(diff),
				sum,
				pn.GetIndex(sum),
			)

			minDiff = diff

			// TODO: Prove that the difference of the first pair of such
			// pentagon numbers is the smallest.

			stop = true
		}
	}

	return "044", minDiff
}
