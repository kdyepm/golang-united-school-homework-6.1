package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (t Triangle) CalcPerimeter() float64 {
	return t.Side * 3
}

// CalcArea returns calculation result of area
func (t Triangle) CalcArea() float64 {
	return math.Sqrt(3) * t.Side * t.Side / 4
}
