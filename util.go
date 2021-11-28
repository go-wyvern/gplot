package gplot

import (
	"fmt"
	"reflect"

	"gonum.org/v1/plot/plotter"
)

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
