package main

import (
	"github.com/go-wyvern/gplot"
)

type index struct{}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:1
	a := []float64{1,2,1.2,1.4}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
	b := []float64{1.4,2.1,1.1,1.2}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:3
	c := []float64{1.5,2.1,1.1,1.2}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:4
	gplot.NominalX("a","b","c")
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:5
	gplot.Bar(a, b, c)
}

func main() {
	gplot.Gopt_Plot_Main(new(index))
}
