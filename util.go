package gplot

import (
	"fmt"
	"math"
	"reflect"

	"gonum.org/v1/plot/plotter"
)

var (
	// const pi 3.1415926
	Pi = math.Pi

	Sin = math.Sin
	Cos = math.Cos
	Tag = math.Tan
	Log = math.Log
)

func Gopt_Figure_Pi() float64 {
	return Pi
}

func Gopt_Figure_Sin(x float64) float64 {
	return Sin(x)
}

func Gopt_Figure_Cos(x float64) float64 {
	return Cos(x)
}

func Gopt_Figure_Log(x float64) float64 {
	return Log(x)
}

func Gopt_Figure_Tag(x float64) float64 {
	return Tag(x)
}

func buildLinePoints(x, y interface{}) plotter.XYs {
	xKind := reflect.TypeOf(x).Kind()
	yKind := reflect.TypeOf(y).Kind()
	if xKind != reflect.Slice || yKind != reflect.Slice {
		fmt.Println("x or y are not slice type")
		return plotter.XYs{}
	}
	xValue := reflect.ValueOf(x)
	yValue := reflect.ValueOf(y)
	if xValue.Len() != yValue.Len() {
		fmt.Println("The length of the slice x is not equal to the slice y")
		return plotter.XYs{}
	}
	xys := plotter.XYs{}
	for i := 0; i < xValue.Len(); i++ {
		xys = append(xys, plotter.XY{
			X: xValue.Index(i).Float(),
			Y: yValue.Index(i).Float(),
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
