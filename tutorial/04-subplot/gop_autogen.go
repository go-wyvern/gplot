package main

import (
	gplot "github.com/go-wyvern/gplot"
	math "math"
)

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:1
	x := gplot.Gopt_Figure_Linspace(this, 0, 2*math.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:2
	y1 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:2
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:2
			_gop_ret = append(_gop_ret, math.Sin(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:2
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:3
	y2 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:3
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:3
			_gop_ret = append(_gop_ret, math.Cos(i))
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:3
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:4
	y3 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:4
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:4
			_gop_ret = append(_gop_ret, i*i)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:4
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:5
	y4 := func() (_gop_ret []float64) {
		for
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:5
		_, i := range x {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:5
			_gop_ret = append(_gop_ret, -1*i*i)
		}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:5
		return
	}()
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:6
	this.Subplot(2, 2, 1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:7
	this.Plot(x, y1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:8
	this.Subplot(2, 2, 2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:9
	this.Plot(x, y2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:10
	this.Subplot(2, 2, 3)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:11
	this.Plot(x, y3)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:12
	this.Subplot(2, 2, 4)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:13
	this.Plot(x, y4)
}
func main() {
	gplot.Gopt_Figure_Main(new(index))
}
