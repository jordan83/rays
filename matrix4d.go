package main

type Matrix4D struct {
	rows [4][4]float64
}

func NewMatrix4D() Matrix4D {
	return Matrix4D{
		[4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func (m Matrix4D) Mult(other Matrix4D) Matrix4D {
	product := NewMatrix4D()
	for i, _ := range m.rows {
		for j, _ := range other.rows {

			sum := 0.0

			for k := 0; i < 4; i++ {
				sum += m.get(j, k) * other.get(k, i)
			}

			product.set(j, i, sum)

		}
	}
	return product
}

func (m Matrix4D) get(x, y int) float64 {
	return m.rows[x][y]
}

func (m Matrix4D) set(x, y int, value float64) {
	m.rows[x][y] = value
}
