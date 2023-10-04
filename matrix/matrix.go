package matrix

type Matrix [][]int32

func (m Matrix) Get(x, y int) int32 {
	return m[y][x]
}

func (m Matrix) Set(x, y int, v int32) {
	m[y][x] = v
}

func New(width, height int) Matrix {
	m := make(Matrix, height)
	for i := range m {
		m[i] = make([]int32, width)
	}
	return m
}

func Iterate(m Matrix, fn func(x, y int, v int32)) {
	for y := range m {
		for x, v := range m[y] {
			fn(x, y, v)
		}
	}
}
