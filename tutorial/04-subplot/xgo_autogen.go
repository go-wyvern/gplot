// Code generated by xgo (XGo); DO NOT EDIT.

package main

import "github.com/go-wyvern/gplot"

const _ = true

type index struct {
	gplot.Figure
}
//line index.plot:1
func (this *index) MainEntry() {
//line index.plot:1:1
	x := gplot.Linspace(0, 2*gplot.Pi, 20)
//line index.plot:2:1
	this.Subplot(2, 2, 1)
//line index.plot:3:1
	this.Plot__2(x, gplot.Sin__1(x))
//line index.plot:4:1
	this.Subplot(2, 2, 2)
//line index.plot:5:1
	this.Plot__2(x, gplot.Cos__1(x))
//line index.plot:6:1
	this.Subplot(2, 2, 3)
//line index.plot:7:1
	this.Plot__2(x, (gplot.Vector).Gop_Mul(x, x))
//line index.plot:8:1
	this.Subplot(2, 2, 4)
//line index.plot:9:1
	this.Plot__2(x, (gplot.Vector).Gop_Mul(gplot.Scale(2, x), x))
}
func (this *index) Main() {
	gplot.Gopt_Figure_Main(this)
}
func main() {
	new(index).Main()
}
