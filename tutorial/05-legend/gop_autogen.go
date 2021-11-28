package main

import gplot "github.com/go-wyvern/gplot"

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/05-legend/index.plot:1
	x := gplot.Linspace(0, 2*gplot.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/05-legend/index.plot:2
	this.Legend("sin(x)", "cos(x)")
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/05-legend/index.plot:3
	this.Plot__2(x, gplot.Sin__1(x), x, gplot.Cos__1(x))
}
func main() {
	gplot.Gopt_Figure_Main(new(index))
}
