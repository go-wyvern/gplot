package gplot

import (
	"fmt"
	"math"
	"reflect"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
)

var (
	// const pi 3.1415926
	Pi = math.Pi

	Sin__0 = math.Sin
	Cos__0 = math.Cos
	Tan__0 = math.Tan
	Log__0 = math.Log
)

func Sin__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Sin(v.AtVec(i)))
	}
	return Vector{vec}
}

func Cos__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Cos(v.AtVec(i)))
	}
	return Vector{vec}
}

func Tan__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Tan(v.AtVec(i)))
	}
	return Vector{vec}
}

func Log__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Log(v.AtVec(i)))
	}
	return Vector{vec}
}

func buildLinePoints(x, y Vector) plotter.XYs {
	if x.Len() != y.Len() {
		fmt.Println("The length of the slice x is not equal to the slice y")
		return plotter.XYs{}
	}
	xys := plotter.XYs{}
	for i := 0; i < x.Len(); i++ {
		xys = append(xys, plotter.XY{
			X: x.AtVec(i),
			Y: y.AtVec(i),
		})
	}
	return xys
}

func buildBarGroup(g interface{}) plotter.Values {
	gKind := reflect.TypeOf(g).Kind()
	if gKind != reflect.Slice {
		return plotter.Values{}
	}
	gValue := reflect.ValueOf(g)
	values := plotter.Values{}
	for i := 0; i < gValue.Len(); i++ {
		values = append(values, gValue.Index(i).Float())
	}
	return values
}
