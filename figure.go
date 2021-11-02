package gplot

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type Figure struct {
	Align      [][]*Axis
	pos        Positon
	w, h       float64
	rows, cols int
	canvas     *vgimg.Canvas
}

type Positon struct {
	row, col int
}

var defaultFigure = &Figure{
	w: float64(4 * vg.Inch),
	h: float64(3 * vg.Inch),
	pos: Positon{
		row: 0,
		col: 0,
	},
	rows: 1,
	cols: 1,
	Align: [][]*Axis{
		{
			NewAxis("Title", "X", "Y"),
		},
	},
}

func (f *Figure) Xlabel(xlabel string) {
	f.Align[f.pos.row][f.pos.col].plot.X.Label.Text = xlabel
}

func (f *Figure) Ylabel(ylabel string) {
	f.Align[f.pos.row][f.pos.col].plot.Y.Label.Text = ylabel
}

func (f *Figure) Title(title string) {
	f.Align[f.pos.row][f.pos.col].plot.Title.Text = title
}

func (f *Figure) NominalX(names ...string) {
	f.Align[f.pos.row][f.pos.col].nominalX = names
}

func (f *Figure) Legend(labels ...string) {
	f.Align[f.pos.row][f.pos.col].legend = labels
}

func (f *Figure) Plot(args ...interface{}) {
	f.addLinePoints(args...)    

	fmt.Println(f.pos.col,f.pos.row,f.cols,f.rows)
	if f.pos.col+1 == f.cols && f.pos.row+1 == f.rows {
		f.show(f.draw())
	}
}

func (f *Figure) Bar(args ...interface{}) {
	f.addBarPoints(args...)

	if f.pos.col+1 == f.cols && f.pos.row+1 == f.rows {
		f.show(f.draw())
	}
}

func (f *Figure) addLinePoints(args ...interface{}) {
	axis := f.Align[f.pos.row][f.pos.col]
	if len(args)%2 == 0 {
		for i := 0; i < len(args)/2; i++ {
			x := args[2*i]
			y := args[2*i+1]
			points := buildLinePoints(x, y)
			line, err := plotter.NewLine(points)
			if err != nil {
				log.Fatal(err)
			}

			line.Color = plotutil.Color(i)
			if len(axis.legend) > 0 {
				axis.plot.Legend.Add(axis.legend[i], line)
			}
			axis.plot.Add(line)
			// plotutil.AddLinePoints(f.Align[f.pos.row][f.pos.col].plot, points)
		}
	}
}

func (f *Figure) addBarPoints(args ...interface{}) {
	axis := f.Align[f.pos.row][f.pos.col]
	l := len(args)
	for i, arg := range args {
		w := vg.Points(16)
		group := buildBarGroup(arg)
		bar, _ := plotter.NewBarChart(group, w)

		bar.Color = plotutil.Color(i)
		bar.Offset = -1*vg.Length(l-1)*w/2 + w*vg.Length(i)

		if len(axis.legend) > 0 {
				axis.plot.Legend.Add(axis.legend[i], bar)
		}
		axis.plot.Add(bar)
	}
	if len(axis.nominalX) > 0 {
		axis.plot.NominalX(axis.nominalX...)
	}
}

func (f *Figure) draw() image.Image {
	w, h := vg.Length(f.w*float64(f.cols)), vg.Length(f.h*float64(f.rows))
	convas := vgimg.NewWith(vgimg.UseWH(w, h), vgimg.UseBackgroundColor(color.Gray16{0xaaaa}))
	dc := draw.New(convas)
	t := draw.Tiles{
		Rows:      f.rows,
		Cols:      f.cols,
		PadTop:    vg.Length(20),
		PadBottom: vg.Length(20),
		PadRight:  vg.Length(20),
		PadLeft:   vg.Length(20),
		PadX:      vg.Length(20),
		PadY:      vg.Length(20),
	}
	plots := make([][]*plot.Plot, f.rows)
	for j := 0; j < f.rows; j++ {
		plots[j] = make([]*plot.Plot, f.cols)
	}
	canvases := plot.Align(plots, t, dc)
	for j := 0; j < f.rows; j++ {
		for i := 0; i < f.cols; i++ {
			axis := f.Align[j][i]
			if axis != nil {
				plots[j][i] = axis.plot
				if plots[j][i] != nil {
					plots[j][i].Draw(canvases[j][i])
				}
			}
		}
	}
	return convas.Image()
}

func (f *Figure) show(img image.Image) {
	image := canvas.NewImageFromImage(img)
	image.FillMode = canvas.ImageFillOriginal
	app := app.New()
	w := app.NewWindow("Figure")
	w.SetContent(image)
	w.SetPadded(false)
	w.ShowAndRun()
}

func (f *Figure) Subplot(x, y, pos int) {
	f.rows = x
	f.cols = y
	var col int
	row := pos / y
	if pos%y == 0 {
		col = y - 1
		row = row - 1
	} else {
		col = pos%y - 1
	}

	f.pos.row = row
	f.pos.col = col

	rows := len(f.Align)
	if rows != x {
		f.Align = make([][]*Axis, x)
	}
	for j := 0; j < rows; j++ {
		if len(f.Align[j]) != y {
			f.Align[j] = make([]*Axis, y)
		}
	}

	if f.Align[row][col] == nil {
		f.Align[row][col] = NewAxis(fmt.Sprintf("title-%d-%d", row, col), "X", "Y")
	}
}
