package main

import (
	"math"
)

// Shapes abstracted through an interface
type Shape interface {
	SurfaceArea() float64
}

type Square struct {
	Side float64
}

func (s Square) SurfaceArea() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) SurfaceArea() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Width, Height float64
}

func (t Triangle) SurfaceArea() float64 {
	return 0.5 * t.Width * t.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) SurfaceArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Shapes Union with type enum
type ShapeType int

const (
	Invalid ShapeType = iota
	SquareType
	RectangleType
	TriangleType
	CircleType
)

type ShapeUnion struct {
	Type   ShapeType
	Width  float64 // Width doubles for Radius on circles and Side of a Square
	Height float64
}

func (su ShapeUnion) SurfaceArea() float64 {
	switch su.Type {
	case SquareType:
		return su.Width * su.Width
	case RectangleType:
		return su.Width * su.Height
	case TriangleType:
		return 0.5 * su.Width * su.Height
	case CircleType:
		return math.Pi * su.Width * su.Width
	default:
		panic("No valid type")
	}
}

var ctable []float64 = []float64{0.0, 1.0, 1.0, 0.5, math.Pi}

/* this depends on the type enum matching
 * I tried a map first, but that resulted in awefull scores
 * */

func SurfaceAreaSwitchTable(su ShapeUnion) float64 {
	return ctable[su.Type] * su.Width * su.Height
}

// List of shapes
const amount int = 10000

var Shapes [amount]Shape
var ShapeUnions [amount]ShapeUnion

func init() {
	// Populate the Shapes and ShapeUnions slices
	for i := 0; i < amount; i++ {
		switch i % 4 {
		case 0:
			var side = float64(((i * i) % 100) + 1)
			Shapes[i] = Square{Side: side}
			//Shapes = append(Shapes, Square{Side: side})
			//ShapeUnions = append(ShapeUnions,
			ShapeUnions[i] = ShapeUnion{
				Type:   SquareType,
				Width:  side,
				Height: side,
			}
		case 1:
			var width = float64(((i * i) % 100) + 1)
			var height = float64(((i * i) % 100) + 1)
			Shapes[i] = Rectangle{
				Width:  width,
				Height: height,
			}
			ShapeUnions[i] = ShapeUnion{
				Type:   RectangleType,
				Width:  width,
				Height: height,
			}
			// Shapes = append(Shapes, Rectangle{
			// 	Width:  width,
			// 	Height: height,
			// })
			// ShapeUnions = append(ShapeUnions, ShapeUnion{
			// 	Type:   RectangleType,
			// 	Width:  width,
			// 	Height: height,
			// })
		case 2:
			var width = float64(((i * i) % 100) + 1)
			var height = float64(((i * i) % 100) + 1)
			Shapes[i] = Triangle{
				Width:  width,
				Height: height,
			}
			ShapeUnions[i] = ShapeUnion{
				Type:   TriangleType,
				Width:  width,
				Height: height,
			}
			// Shapes = append(Shapes, Triangle{
			// 	Width:  width,
			// 	Height: height,
			// })
			// ShapeUnions = append(ShapeUnions, ShapeUnion{
			// 	Type:   TriangleType,
			// 	Width:  width,
			// 	Height: height,
			// })
		case 3:
			var radius = float64(((i * i) % 100) + 1)
			Shapes[i] = Circle{Radius: radius}
			ShapeUnions[i] = ShapeUnion{
				Type:   CircleType,
				Width:  radius,
				Height: radius,
			}
			// Shapes = append(Shapes, Circle{Radius: radius})
			// ShapeUnions = append(ShapeUnions, ShapeUnion{
			// 	Type:   CircleType,
			// 	Width:  radius,
			// 	Height: radius,
			// })
		}
	}
}
