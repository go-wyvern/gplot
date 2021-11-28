package main

import gplot "github.com/go-wyvern/gplot"

type index struct {
	gplot.Figure
}

func (this *index) MainEntry() {
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/07-bar/index.plot:1
	a := []float64{1, 3, 5, 7.0}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/07-bar/index.plot:2
	b := []float64{1, 3, 5.0, 7}
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/07-bar/index.plot:3
	this.NominalX("s1", "s2", "s3", "s4")
//line /Users/wuxinyi/go/src/github.com/go-wyvern/gplot/tutorial/07-bar/index.plot:4
	this.Bar(a, b)
}
func main() {
	gplot.Gopt_Figure_Main(new(index))
}
