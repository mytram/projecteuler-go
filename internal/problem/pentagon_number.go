package problem

func pentagonNumber(n uint) uint {
	//  Pn=n(3n−1)/2

	return n * (3*n - 1) / 2
}

type pentagonNumbers struct {
	index          uint
	numbers        map[uint]uint
	reverseNumbers map[uint]uint
}

func (this *pentagonNumbers) getByIndex(index uint) uint {
	if this.index >= index {
		return this.numbers[index]
	}

	for i := this.index; i <= index; i++ {
		pentagon := pentagonNumber(i)

		this.index = i
		this.numbers[i] = pentagon
		this.reverseNumbers[pentagon] = i
	}

	return this.numbers[index]
}

func (this *pentagonNumbers) addMoreTill(limit uint) {
	if this.numbers[this.index] >= limit {
		return
	}

	for {
		this.index += 1
		pentagon := pentagonNumber(this.index)

		this.numbers[this.index] = pentagon
		this.reverseNumbers[pentagon] = this.index

		if pentagon >= limit {
			return
		}
	}
}

func (this *pentagonNumbers) found(number uint) bool {
	this.addMoreTill(number)
	_, ok := this.reverseNumbers[number]
	return ok
}

func newPentagonNumbers() *pentagonNumbers {
	pn := &pentagonNumbers{
		index:          1,
		numbers:        make(map[uint]uint),
		reverseNumbers: make(map[uint]uint),
	}

	pn.numbers[1] = 1

	return pn
}
