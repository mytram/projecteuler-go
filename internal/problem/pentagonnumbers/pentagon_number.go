package pentagonnumbers

type pentagonNumbers struct {
	index          uint
	numbers        map[uint]uint
	reverseNumbers map[uint]uint
}

func New() *pentagonNumbers {
	pn := &pentagonNumbers{
		index:          1,
		numbers:        make(map[uint]uint),
		reverseNumbers: make(map[uint]uint),
	}

	pn.numbers[1] = 1

	return pn
}

func (this *pentagonNumbers) Get(index uint) uint {
	if this.index >= index {
		return this.numbers[index]
	}

	for i := this.index; i <= index; i++ {
		this.addNext()
	}

	return this.numbers[index]
}

func (this *pentagonNumbers) GetIndex(number uint) uint {
	this.addMoreTill(number)
	return this.reverseNumbers[number]
}

func (this *pentagonNumbers) IsPentagonNumber(number uint) bool {
	this.addMoreTill(number)
	_, ok := this.reverseNumbers[number]
	return ok
}

func (this *pentagonNumbers) addNext() uint {
	i := this.index + 1
	pentagon := i * (3*i - 1) / 2
	this.index = i
	this.numbers[i] = pentagon
	this.reverseNumbers[pentagon] = i

	return pentagon
}

func (this *pentagonNumbers) addMoreTill(limit uint) {
	for pentagon := this.Get(this.index); pentagon < limit; {
		pentagon = this.addNext()
	}
}
