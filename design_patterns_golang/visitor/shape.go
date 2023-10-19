package main

type Shape interface {
	getType() string
	accept(Visitor)
}

// square

type Square struct {
	s float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// rectangle

type Rectangle struct {
	l float64
	b float64
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

// circle

type Circle struct {
	r float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
