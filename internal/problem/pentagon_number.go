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

func (this *pentagonNumbers) addNext() uint {
	i := this.index + 1
	pentagon := pentagonNumber(i)
	this.index = i
	this.numbers[i] = pentagon
	this.reverseNumbers[pentagon] = i

	return pentagon
}

func (this *pentagonNumbers) get(index uint) uint {
	if this.index >= index {
		return this.numbers[index]
	}

	for i := this.index; i <= index; i++ {
		this.addNext()
	}

	return this.numbers[index]
}

func (this *pentagonNumbers) addMoreTill(limit uint) {
	for pentagon := this.get(this.index); pentagon < limit; {
		pentagon = this.addNext()
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
