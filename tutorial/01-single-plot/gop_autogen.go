package main

import (
	gplot "github.com/go-wyvern/gplot"
	math "math"
)

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:1
	x := gplot.Gopt_Figure_Linspace(this, 0, 2*math.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:2
	y := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:2
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:2
			_gop_ret = append(_gop_ret, math.Sin(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:2
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/01-single-plot/index.plot:3
	this.Plot(x, y)
}
func main() {
	gplot.Gopt_Figure_Main(new(index))
}
