package gplot

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Axis struct {
	p *plot.Plot

	legend []string
}

func NewAxis() *Axis {
	p := plot.New()
	p.Title.Padding = 10
	return &Axis{
		p: p,
	}
}

func (c *Axis) Plot__0(args ...[]float64) {
	var vecs []Vector
	for _, arg := range args {
		vecs = append(vecs, newVec(arg))
	}
	c.addLine(vecs...)
}

func (c *Axis) Plot__1(args ...[]int) {
	var vecs []Vector
	for _, arg := range args {
		i := []float64{}
		for _, a := range arg {
			i = append(i, float64(a))
		}
		vecs = append(vecs, newVec(i))
	}
	c.addLine(vecs...)
}

func (c *Axis) Plot__2(args ...Vector) {
	c.addLine(args...)
}

func (c *Axis) Bar(args ...interface{}) {
	c.addBar(args...)
}

func (a *Axis) Xlabel(xlabel string) {
	a.p.X.Label.Text = xlabel
}

func (a *Axis) Ylabel(ylabel string) {
	a.p.Y.Label.Text = ylabel
}

func (a *Axis) Title(title string) {
	a.p.Title.Text = title
}

func (a *Axis) Legend(labels []string) {
	a.legend = labels
}

func (c *Axis) NominalX(names ...string) {
	c.p.NominalX(names...)
}

func (c *Axis) addLine(args ...Vector) {
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
			c.p.Legend.Add(c.legend[i], line)
		}
		c.p.Add(line)
	}
}

func (c *Axis) addBar(args ...interface{}) {
	l := len(args)
	for i, arg := range args {
		w := vg.Points(16)
		group := buildBarGroup(arg)
		bar, _ := plotter.NewBarChart(group, w)

		bar.Color = plotutil.Color(i)
		bar.Offset = -1*vg.Length(l-1)*w/2 + w*vg.Length(i)

		if len(c.legend) > 0 {
			c.p.Legend.Add(c.legend[i], bar)
		}
		c.p.Add(bar)
	}
}
