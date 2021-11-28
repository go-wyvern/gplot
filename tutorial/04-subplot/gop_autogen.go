package main

import gplot "github.com/go-wyvern/gplot"

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:1
	x := gplot.Linspace(0, 2*gplot.Pi, 20)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:2
	this.Subplot(2, 2, 1)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:3
	this.Plot__2(x, gplot.Sin__1(x))
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:4
	this.Subplot(2, 2, 2)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:5
	this.Plot__2(x, gplot.Cos__1(x))
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:6
	this.Subplot(2, 2, 3)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:7
	this.Plot__2(x, x.Gop_Mul(x))
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:8
	this.Subplot(2, 2, 4)
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/04-subplot/index.plot:9
	this.Plot__2(x, gplot.Scale(2, x).Gop_Mul(x))
}
func main() {
	gplot.Gopt_Figure_Main(new(index))
}
