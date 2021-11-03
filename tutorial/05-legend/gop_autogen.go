package main

import (
	"github.com/go-wyvern/gplot"
)

type index struct{}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:1
	x := gplot.Linspace(0,2*gplot.Pi,20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
	y1 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
			_gop_ret = append(_gop_ret, gplot.Sin(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:2
		return
	}()
	y2 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:3
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:3
			_gop_ret = append(_gop_ret, gplot.Cos(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:3
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:4
	gplot.Legend("sin(x)","cos(x)")
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:5
	gplot.Plot(x, y1, x, y2)
}

func main() {
	gplot.Gopt_Game_Main(new(index))
}