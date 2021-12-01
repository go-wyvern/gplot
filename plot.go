package gplot

import (
	"log"
	"reflect"

	"gonum.org/v1/gonum/floats"
)

const (
	GopPackage = true
	Gop_game   = "Figure"
	Gop_sprite = "Axis"
)

type Ploter interface {
	initPlot()
	finishPlot()
}

// Gopt_Figure_Main is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Main(plot Ploter) {
	plot.initPlot()
	defer plot.finishPlot()
	plot.(interface{ MainEntry() }).MainEntry()
}

// Gopt_Figure_Run is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Run(plot Ploter, x, y int) {
	v := reflect.ValueOf(plot).Elem()
	t := reflect.TypeOf(plot).Elem()
	p := instance(v)
	pos := 0
	for i, n := 0, v.NumField(); i < n; i++ {
		typ := t.Field(i).Type
		m, ok := reflect.PtrTo(typ).MethodByName("Main")
		if ok {
			parent, axis := instanceAxis(t.Field(i))
			pos += 1
			m.Func.Call([]reflect.Value{parent})
			p.Subplot(x, y, pos)
			p.Align[p.pos.row][p.pos.col] = axis
		}
	}
}

func Linspace(l, r float64, n int) Vector {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return newVec(dst)
}

func instance(plotter reflect.Value) *Figure {
	fld := plotter.FieldByName("Figure")
	if !fld.IsValid() {
		log.Panicf("type %v doesn't has field gplot.Figure", plotter.Type())
	}
	return fld.Addr().Interface().(*Figure)
}

func instanceAxis(field reflect.StructField) (reflect.Value, *Axis) {
	typ := field.Type
	parent := reflect.New(typ)
	axis := parent.Elem().FieldByName("Axis")
	axis.Set(reflect.ValueOf(NewAxis()).Elem())
	return parent, axis.Addr().Interface().(*Axis)
}
