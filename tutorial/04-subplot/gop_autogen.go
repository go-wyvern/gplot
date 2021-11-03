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
	y3 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:4
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:4
			_gop_ret = append(_gop_ret, i*i)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:4
		return
	}()
	y4 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:5
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:5
			_gop_ret = append(_gop_ret, -1*i*i)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:5
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:6
	gplot.Subplot(2, 2, 1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:7
	gplot.Plot(x, y1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:8
	gplot.Subplot(2, 2, 2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:9
	gplot.Plot(x, y2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:10
	gplot.Subplot(2, 2, 3)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:11
	gplot.Plot(x, y3)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:12
	gplot.Subplot(2, 2, 4)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01_single_plot/plot.p:13
	gplot.Plot(x, y4)
}

func main() {
	gplot.Gopt_Game_Main(new(index))
}
