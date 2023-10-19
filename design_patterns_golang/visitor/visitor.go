package main

import "fmt"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// Area Calculator

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	var area float64 = s.s * s.s
	fmt.Printf("area for square: %f\n", area)
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	var area float64 = 22 / 7 * s.r * s.r
	fmt.Printf("area for circle: %f\n", area)
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	var area float64 = s.l * s.b
	fmt.Printf("area for rectangle: %f\n", area)
}

// Perimeter Calculator

type PerimeterCalculator struct {
	perimeter int
}

func (a *PerimeterCalculator) visitForSquare(s *Square) {
	var peri float64 = 4 * s.s
	fmt.Printf("perimeter for square: %f\n", peri)
}

func (a *PerimeterCalculator) visitForCircle(s *Circle) {
	var peri float64 = 2 * 3.14 * s.r
	fmt.Printf("perimeter for circle: %f\n", peri)
}

func (a *PerimeterCalculator) visitForRectangle(s *Rectangle) {
	var peri float64 = s.l * s.b
	fmt.Printf("perimeter for rectangle: %f\n", peri)
}

//area for square: 16.000000
//area for rectangle: 6.000000
//area for circle: 147.000000
//perimeter for square: 16.000000
//perimeter for rectangle: 6.000000
//perimeter for circle: 43.960000
