package gplot

import (
	"log"
	"reflect"

	"gonum.org/v1/gonum/floats"
)

type Ploter interface {
	initPlot()
	finishPlot()
}

// Gopt_Plot_Main is required by Go+ compiler as the entry of a .plot project.
func Gopt_Plot_Main(plot Ploter) {
	plot.initPlot()
	defer plot.finishPlot()
	plot.(interface{ MainEntry() }).MainEntry()
}

// gop enter func
func Plot(plot Ploter, args ...interface{}) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Plot(args...)
}

func Bar(plot Ploter, args ...interface{}) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Bar(args...)
}

func Subplot(plot Ploter, row, col, pos int) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Subplot(row, col, pos)
}

func Legend(plot Ploter, labels ...string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Legend(labels...)
}

func Xlabel(plot Ploter, xlabel string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Xlabel(xlabel)
}

func Ylabel(plot Ploter, ylabel string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Ylabel(ylabel)
}

func Title(plot Ploter, title string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Title(title)
}

func NominalX(plot Ploter, names ...string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.NominalX(names...)
}

func Linspace(plot Ploter, l, r float64, n int) []float64 {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return dst
}

func instance(plotter reflect.Value) *Figure {
	fld := plotter.FieldByName("Figure")
	if !fld.IsValid() {
		log.Panicf("type %v doesn't has field gplot.Figure", plotter.Type())
	}
	return fld.Addr().Interface().(*Figure)
}
