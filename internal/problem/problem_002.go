package problem

func Solve002() (string, int64) {
	sum := int64(0)

	f0, f1 := 0, 1

	for fn := 1; fn < 4_000_000; fn = f0 + f1 {
		if fn%2 == 0 {
			sum += int64(fn)
		}

		f0, f1 = f1, fn
	}

	return "002", sum
}
