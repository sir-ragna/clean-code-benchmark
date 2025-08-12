
Some simple Golang benchmarks.


# Inspiration was Casey Muratori

Casey Muratori has a series of lessons that focus on performance.
One of his free bonus video's and blogposts he dives into the performance impact of some typical **clean code** rules.

https://www.computerenhance.com/p/clean-code-horrible-performance

I highly recommend the included video.

Anyway, he benchmarks the surface calculation of shapes in multiple different ways.

 - through abstracted inheretance
 - a variation with an unrolled loop
 - a struct union and switch table
 - based on the previous patterns, he creates a table that substitues one parameter in a generalized formula

The performance gains are striking of course.

Other people have experimente with rewritting it in other languages and benchmarking it, such as rust.
https://www.reddit.com/r/rust/comments/11fkfib/i_ported_casey_muratoris_c_example_of_clean_code/

# My attempt in Golang

Golang doesn't have inheritance and virtual function. At least not as far as I know. I haven't kept pace with the features of the last 10 version or so.

What Golang does have are interfaces. I wonder whether the performance impact of interfaces in Go is similar to the impact of inheritance and virtual functions in C++.

Hence, I implemented my own version.

## interface code

I start with a Shape interface.

```go
type Shape interface {
	SurfaceArea() float64
}
```

And then I implement this interface for the shapes **rectangle**, **square**, **circle** and **triangle**.

```go
type Square struct {
	Side float64
}

func (s Square) SurfaceArea() float64 {
	return s.Side * s.Side
}
```

To be able to benchmark this, I create an array of 10000 shapes.

```go
const amount int = 10000
var Shapes [amount]Shape
```

I populate these in the `init()` function.
This gets called when this module is loaded.

I used to generate them randomly, but I wanted them to be deterministic, so I could test the correctness of my code.

That is why I used modulo to generate values based on the index.

```go 
func init() {
	for i := 0; i < amount; i++ {
		switch i % 4 {
		case 0:
			var side = float64(((i * i) % 100) + 1)
			Shapes[i] = Square{Side: side}
		case 1:
			var width = float64(((i * i) % 100) + 1)
			var height = float64(((i * i) % 100) + 1)
			Shapes[i] = Rectangle{Width:  width, Height: height}
		case 2:
			var width = float64(((i * i) % 100) + 1)
			var height = float64(((i * i) % 100) + 1)
			Shapes[i] = Triangle{Width:  width, Height: height}
		case 3:
			var radius = float64(((i * i) % 100) + 1)
			Shapes[i] = Circle{Radius: radius}
		}
	}
}
```

The code to calculate the surface is very straightforward.
This is supposed to be the _clean code_ example after all.

```go
var sum float64
for _, shape := range Shapes {
    sum += shape.SurfaceArea()
}
```

Just like Casey, I made a variation to test what the effect is of unrolling the loop of the calculation.

```go
var accum1, accum2, accum3, accum4, accum5, accum6, accum7, accum8, sum float64

n := len(Shapes)
i := 0

// Unroll the loop by 8
for ; i+8 < n; i += 8 {
    accum1 += Shapes[i].SurfaceArea()
    accum2 += Shapes[i+1].SurfaceArea()
    accum3 += Shapes[i+2].SurfaceArea()
    accum4 += Shapes[i+3].SurfaceArea()
    accum5 += Shapes[i+4].SurfaceArea()
    accum6 += Shapes[i+5].SurfaceArea()
    accum7 += Shapes[i+6].SurfaceArea()
    accum8 += Shapes[i+7].SurfaceArea()
}
sum = accum1 + accum2 + accum3 + accum4 + accum5 + accum6 + accum7 + accum8
// Handle remaining elements
for ; i < n; i++ {
    sum += Shapes[i].SurfaceArea()
}
```

This does two things.

1) Less branches, although good compiler will already do this type of optimization.

2) By splitting up the accumulators, the instructions are not dependant on eachother. 
If the instructions are not dependant on eachother for execution, the CPU can keep executing them without having to wait on the previous instruction to finish.

## Using the same stuct with a Type

Instead of using an interface, we use a struct and a Switch statement to perform the calculation.

```go
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
```

# How to run the different tests

Run `go test -bench '.*'`.

```
goos: windows
goarch: amd64
pkg: sources/clean-code-bench
cpu: 12th Gen Intel(R) Core(TM) i5-1245U
BenchmarkShapesSurfaceAreaDynamicDispatch-12                       53188             22042 ns/op
BenchmarkShapesSurfaceAreaDynamicDispatchUnrolled-12               99675             11541 ns/op
BenchmarkShapesSurfaceAreaUnionSwitch-12                           51031             23798 ns/op
BenchmarkShapesSurfaceAreaUnionSwitchUnrolled-12                   71772             16470 ns/op
BenchmarkShapeAreaFunctionWithctable-12                            51168             23414 ns/op
BenchmarkShapeAreaFunctionWithctableUnrolled-12                    78050             14901 ns/op
PASS
ok      sources/clean-code-bench        8.103s
```

The first column is the amount of iteration done. Higher is better, the second one is time per operation where lower is better.

It seems that the unrolled loop version of the surface calculation that uses the interface wins here. Even over the struct switch table and lookup table optimizations.

