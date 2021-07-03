package problem

import "fmt"

func Solve044() (string, uint) {
	minDiff := ^uint(0)
	stop := false

	pn := newPentagonNumbers()

	n := uint(0)

	for !stop {
		n += 1
		pentagon := pn.getByIndex(n)

		for i := n - 1; i > 0; i-- {
			diff := pentagon - pn.numbers[i]

			if diff >= minDiff {
				break
			}

			if !pn.found(diff) {
				continue
			}

			sum := pentagon + pn.numbers[i]

			if !pn.found(sum) {
				continue
			}

			fmt.Printf("n=%d pentagon=%d i=%d pentagon=%d diff=%d diffN=%d sum=%d sumN=%d\n",
				n, pentagon, i,
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
