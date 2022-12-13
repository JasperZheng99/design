package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	triangle := &Triangle{}
	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	triangle.accept(areaCalculator)

	fmt.Println("访问者模式")
	middleCoodrdinates := &MiddleCoordinates{}
	square.accept(middleCoodrdinates)
	circle.accept(middleCoodrdinates)
	triangle.accept(middleCoodrdinates)

}
