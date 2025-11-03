package main

import "fmt"

type Point struct {
	x float64
	y float64
}

type Line struct {
	initialPoint Point
	endingPoint  Point
}

func (l Line) lineInclination() float64 {
	numerator := l.endingPoint.y - l.initialPoint.y
	denominator := l.endingPoint.x - l.initialPoint.x

	return numerator / denominator
}

func goStruct() {
	myLine := Line{initialPoint: Point{x: 0, y: 0}, endingPoint: Point{x: 10, y: 10}}

	fmt.Println("Line inclination", myLine.lineInclination())
}
