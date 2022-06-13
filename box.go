package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity < len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("no capacity left in the current box")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		return b.shapes[i], nil
	} else {
		return nil, fmt.Errorf("wring index. no such an item in the box")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return s, err
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return s, err
	}
	b.shapes[i] = shape
	return s, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, i := range b.shapes {
		sum += i.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	area := 0.0
	for _, i := range b.shapes {
		area += i.CalcArea()
	}
	return area

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	s := []Shape{}
	for i, item := range b.shapes {
		switch item.(type) {
		case Circle:
			extracted, err := b.ExtractByIndex(i)
			if err == nil {
				s = append(s, extracted)
			}
		default:
			continue
		}
	}
	if len(s) == 0 {
		return fmt.Errorf("no circles in the box")
	}

	return nil

}
