package gplot

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Axis struct {
	plot *plot.Plot
	X    string
	Y    string
	Tit  string

	pos      int
	nominalX []string
	legend   []string
}

func NewAxis(title, x, y string) *Axis {
	p := plot.New()
	p.Title.Padding = 10
	return &Axis{
		plot: p,
		X:    x,
		Y:    y,
		Tit:  title,
	}
}

func (c *Axis) Setpos(pos int) {
	c.pos = pos
}

func (c *Axis) Addline__0(args ...[]float64) {
	var vecs []Vector
	for _, arg := range args {
		vecs = append(vecs, newVec(arg))
	}
	c.newline(vecs...)
}

func (c *Axis) Addline__1(args ...[]int) {
	var vecs []Vector
	for _, arg := range args {
		i := []float64{}
		for _, a := range arg {
			i = append(i, float64(a))
		}
		vecs = append(vecs, newVec(i))
	}
	c.newline(vecs...)
}

func (c *Axis) Addline__2(args ...Vector) {
	c.newline(args...)
}

func (c *Axis) newline(args ...Vector) {
	for i := 0; i < len(args)/2; i++ {
		x := args[2*i]
		y := args[2*i+1]
		points := buildLinePoints(x, y)
		line, err := plotter.NewLine(points)
		if err != nil {
			log.Fatal(err)
		}

		line.Color = plotutil.Color(i)
		if len(c.legend) > 0 {
			c.plot.Legend.Add(c.legend[i], line)
		}
		c.plot.Add(line)
	}
}

func (a *Axis) Xlabel(xlabel string) {
	a.plot.X.Label.Text = xlabel
}

func (a *Axis) Ylabel(ylabel string) {
	a.plot.Y.Label.Text = ylabel
}

func (a *Axis) Title(title string) {
	a.plot.Title.Text = title
}

func (a *Axis) Legend(labels []string) {
	a.legend = labels
}
