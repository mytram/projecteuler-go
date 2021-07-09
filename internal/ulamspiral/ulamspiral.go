package ulamspiral

type spiral struct {
	Length int
	n      int
}

type nodeProc func(n, x, y int)

type iter int

func (t iter) times(proc func()) {
	for i := 0; iter(i) < t; i++ {
		proc()
	}
}

func New() *spiral {
	return &spiral{1, 1}
}

func (s *spiral) Circle(proc nodeProc) {
	x := (s.Length - 1) / 2
	y := -x

	s.Length += 2

	steps := iter(s.Length - 1)

	s.n, x, y = r(s.n, x, y, 1, proc)
	s.n, x, y = u(s.n, x, y, iter(s.Length-2), proc)
	s.n, x, y = l(s.n, x, y, steps, proc)
	s.n, x, y = d(s.n, x, y, steps, proc)
	s.n, x, y = r(s.n, x, y, steps, proc)
}

func r(n, x, y int, steps iter, proc nodeProc) (int, int, int) {
	steps.times(func() {
		n, x = n+1, x+1
		proc(n, x, y)
	})

	return n, x, y
}

func l(n, x, y int, steps iter, proc nodeProc) (int, int, int) {
	steps.times(func() {
		n, x = n+1, x-1
		proc(n, x, y)
	})

	return n, x, y
}

func u(n, x, y int, steps iter, proc nodeProc) (int, int, int) {
	steps.times(func() {
		n, y = n+1, y+1
		proc(n, x, y)
	})

	return n, x, y
}

func d(n, x, y int, steps iter, proc nodeProc) (int, int, int) {
	steps.times(func() {
		n, y = n+1, y-1
		proc(n, x, y)
	})

	return n, x, y
}
