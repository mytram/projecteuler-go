package problem

import "fmt"

func Solve044() (string, uint) {
	minDiff := ^uint(0)
	stop := false

	pn := newPentagonNumbers()

	for n := uint(1); !stop; n++ {
		nth := pn.get(n)

		for i := n - 1; i > 0; i-- {
			ith := pn.get(i)
			diff := nth - ith

			if diff >= minDiff {
				break
			}

			if !pn.found(diff) {
				continue
			}

			sum := nth + ith

			if !pn.found(sum) {
				continue
			}

			fmt.Printf("n=%d nth=%d i=%d nth=%d diff=%d diffN=%d sum=%d sumN=%d\n",
				n, nth, i,
				pn.numbers[i],
				diff,
				pn.reverseNumbers[diff],
				sum,
				pn.reverseNumbers[sum],
			)

			minDiff = diff

			// TODO: Prove that the difference of the first pair of such
			// pentagon numbers is the smallest.

			stop = true
		}
	}

	return "044", minDiff
}
