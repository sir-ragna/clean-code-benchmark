package main

import (
	"math"
	"testing"
)

var totalsum float64

func init() {
	for _, shape := range Shapes {
		totalsum += shape.SurfaceArea()
	}
	// fmt.Printf("total sum %f\n", totalsum)
}

/* Due to subtle changes in binary representation, a simple equality operator can fail when comparing floats
 * To resolve this, you subtract both numbers and check whether they are small enough.
 */
func floatAreEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-7
}

func BenchmarkShapesSurfaceAreaDynamicDispatch(b *testing.B) {
	for b.Loop() {
		var sum float64
		for _, shape := range Shapes {
			sum += shape.SurfaceArea()
		}
		//fmt.Printf("sum %f\n", sum)
		if !floatAreEqual(sum, totalsum) {
			b.FailNow()
		}
	}
}

func BenchmarkShapesSurfaceAreaDynamicDispatchUnrolled(b *testing.B) {
	for b.Loop() {
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

		if !floatAreEqual(sum, totalsum) {
			//fmt.Printf("unrolled sum %f while totalsum is %f, the difference is %f\n", sum, totalsum, totalsum-sum)
			b.FailNow()
		}
	}
}

func BenchmarkShapesSurfaceAreaUnionSwitch(b *testing.B) {
	for b.Loop() {
		var sum float64
		for _, shape := range ShapeUnions {
			sum += shape.SurfaceArea()
		}
		// fmt.Printf("sum %f\n", sum)
		if !floatAreEqual(sum, totalsum) {
			b.FailNow()
		}
	}
}

func BenchmarkShapesSurfaceAreaUnionSwitchUnrolled(b *testing.B) {
	// b.ReportAllocs() // for -benchmem tests
	for b.Loop() {
		var accum1, accum2, accum3, accum4, accum5, accum6, accum7, accum8, sum float64

		n := len(ShapeUnions)
		i := 0

		// Unroll the loop by 8
		for ; i+8 < n; i += 8 {
			accum1 += ShapeUnions[i].SurfaceArea()
			accum2 += ShapeUnions[i+1].SurfaceArea()
			accum3 += ShapeUnions[i+2].SurfaceArea()
			accum4 += ShapeUnions[i+3].SurfaceArea()
			accum5 += ShapeUnions[i+4].SurfaceArea()
			accum6 += ShapeUnions[i+5].SurfaceArea()
			accum7 += ShapeUnions[i+6].SurfaceArea()
			accum8 += ShapeUnions[i+7].SurfaceArea()
		}
		sum = accum1 + accum2 + accum3 + accum4 + accum5 + accum6 + accum7 + accum8
		// Handle remaining elements
		for ; i < n; i++ {
			sum += ShapeUnions[i].SurfaceArea()
		}

		// fmt.Printf("sum %f\n", sum)
		if !floatAreEqual(sum, totalsum) {
			b.FailNow()
		}
	}
}

func BenchmarkShapeAreaFunctionWithctable(b *testing.B) {
	for b.Loop() {
		var sum float64
		for _, shape := range ShapeUnions {
			sum += SurfaceAreaSwitchTable(shape)
		}
		// fmt.Printf("sum %f\n", sum)
		if !floatAreEqual(sum, totalsum) {
			b.FailNow()
		}
	}
}

func BenchmarkShapeAreaFunctionWithctableUnrolled(b *testing.B) {
	for b.Loop() {
		var accum1, accum2, accum3, accum4, accum5, accum6, accum7, accum8, sum float64

		n := len(ShapeUnions)
		i := 0

		// Unroll the loop by 8
		for ; i+8 < n; i += 8 {
			accum1 += SurfaceAreaSwitchTable(ShapeUnions[i])
			accum2 += SurfaceAreaSwitchTable(ShapeUnions[i+1])
			accum3 += SurfaceAreaSwitchTable(ShapeUnions[i+2])
			accum4 += SurfaceAreaSwitchTable(ShapeUnions[i+3])
			accum5 += SurfaceAreaSwitchTable(ShapeUnions[i+4])
			accum6 += SurfaceAreaSwitchTable(ShapeUnions[i+5])
			accum7 += SurfaceAreaSwitchTable(ShapeUnions[i+6])
			accum8 += SurfaceAreaSwitchTable(ShapeUnions[i+7])
		}
		sum = accum1 + accum2 + accum3 + accum4 + accum5 + accum6 + accum7 + accum8
		// Handle remaining elements
		for ; i < n; i++ {
			sum += ShapeUnions[i].SurfaceArea()
		}
		// fmt.Printf("sum %f\n", sum)
		if !floatAreEqual(sum, totalsum) {
			b.FailNow()
		}
	}
}
