package main

/*

Visitor is a behavioral design pattern that allows
adding new behaviors to existing class hierarchy
without altering any existing code.

*/

func main() {
	square := &Square{s: 4}
	rectangle := &Rectangle{b: 2, l: 3}
	circle := &Circle{r: 7}

	areaCal := &AreaCalculator{}
	// provide area calculator to all shapes
	square.accept(areaCal)
	rectangle.accept(areaCal)
	circle.accept(areaCal)
	//	provide perimeter calculator to all shapes
	periCal := &PerimeterCalculator{}
	square.accept(periCal)
	rectangle.accept(periCal)
	circle.accept(periCal)
}
