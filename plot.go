package gplot

import "gonum.org/v1/gonum/floats"

type Ploter interface{}

// Gopt_Plot_Main is required by Go+ compiler as the entry of a .plot project.
func Gopt_Plot_Main(plot Ploter) {
	// plot.initPlot() TODO
	plot.(interface{ MainEntry() }).MainEntry()
}

// gop enter func
func Plot(args ...interface{}) {
	defaultFigure.Plot(args...)
}

func Bar(args ...interface{}) {
	defaultFigure.Bar(args...)
}

func Subplot(row, col, pos int) {
	defaultFigure.Subplot(row, col, pos)
}

func Legend(labels ...string) {
	defaultFigure.Legend(labels...)
}

func Xlabel(xlabel string) {
	defaultFigure.Xlabel(xlabel)
}

func Ylabel(ylabel string) {
	defaultFigure.Ylabel(ylabel)
}

func Title(title string) {
	defaultFigure.Title(title)
}

func NominalX(names ...string) {
	defaultFigure.NominalX(names...)
}

func Linspace(l, r float64, n int) []float64 {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return dst
}
