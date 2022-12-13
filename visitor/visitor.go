package visitor

import "fmt"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Triangle)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) getType() string {

	return "Square"
}

func (s *Square) accept(visitor Visitor) {
	visitor.visitForSquare(s)
}

type Circle struct {
	radius int
}

func (s *Circle) getType() string {
	return "Circle"
}

func (s *Circle) accept(visitor Visitor) {
	visitor.visitForCircle(s)
}

type Triangle struct {
	l int
	b int
}

func (s *Triangle) getType() string {
	return "Triangle"
}

func (s *Triangle) accept(visitor Visitor) {
	visitor.visitForTriangle(s)
}

// AreaCalculator 具体访问者
type AreaCalculator struct {
	area int
}

func (visitor *AreaCalculator) visitForSquare(square *Square) {
	fmt.Println("Calculating area for square")
}

func (visitor *AreaCalculator) visitForCircle(circle *Circle) {
	fmt.Println("Calculating area for circle")
}

func (visitor *AreaCalculator) visitForTriangle(triangle *Triangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
	func (a *MiddleCoordinates) visitForTriangle(t *Triangle) {
	fmt.Println("Calculating middle point coordinates for Triangle")
}
