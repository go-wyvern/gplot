package gplot

import "gonum.org/v1/plot"

type Axis struct {
	plot  *plot.Plot
	X     string
	Y     string
	Title string

	nominalX []string
	legend   []string
}

func NewAxis(title, x, y string) *Axis {
	p := plot.New()
	p.Title.Padding = 10
	return &Axis{
		plot:  p,
		X:     x,
		Y:     y,
		Title: title,
	}
}
