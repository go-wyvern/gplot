package main

import (
	"github.com/go-wyvern/gplot"
)

type index struct{}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:1
	x := gplot.Linspace(0,2*gplot.Pi,20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
	y := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
			_gop_ret = append(_gop_ret, gplot.Sin(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:3
	gplot.Xlabel("x")
	gplot.Ylabel("sin(x)")
	gplot.Plot(x, y)
}

func main() {
	p:= index{}
	p.MainEntry()
}
