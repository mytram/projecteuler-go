package problem

func Solve001() (string, int) {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
			continue
		}

		if i%5 == 0 {
			sum += i
		}
	}

	return "001", sum
}
