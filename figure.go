package gplot

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type Figure struct {
	Align      [][]*Axis
	pos        Positon
	w, h       float64
	rows, cols int
	app        fyne.App
}

type Positon struct {
	row, col int
}

func (f *Figure) initPlot() {
	f.w = float64(4 * vg.Inch)
	f.h = float64(3 * vg.Inch)
	f.pos = Positon{
		row: 0,
		col: 0,
	}
	f.rows = 1
	f.cols = 1
	f.Align = [][]*Axis{
		{
			NewAxis(),
		},
	}
	f.app = app.New()
}

func (f *Figure) isFinish() bool {
	return f.pos.col+1 == f.cols && f.pos.row+1 == f.rows
}

func (f *Figure) finishPlot() {
	if f.isFinish() {
		f.show(f.draw())
	}
}

func (f *Figure) Xlabel(xlabel string) {
	f.Align[f.pos.row][f.pos.col].Xlabel(xlabel)
}

func (f *Figure) Ylabel(ylabel string) {
	f.Align[f.pos.row][f.pos.col].Ylabel(ylabel)
}

func (f *Figure) Title(title string) {
	f.Align[f.pos.row][f.pos.col].Title(title)
}

func (f *Figure) NominalX(names ...string) {
	f.Align[f.pos.row][f.pos.col].NominalX(names...)
}

func (f *Figure) Legend(labels ...string) {
	f.Align[f.pos.row][f.pos.col].Legend(labels)
}

func (f *Figure) Plot__0(args ...[]float64) {
	f.Align[f.pos.row][f.pos.col].Plot__0(args...)
}

func (f *Figure) Plot__1(args ...[]int) {
	f.Align[f.pos.row][f.pos.col].Plot__1(args...)
}

func (f *Figure) Plot__2(args ...Vector) {
	f.Align[f.pos.row][f.pos.col].Plot__2(args...)
}

func (f *Figure) Bar(args ...interface{}) {
	f.Align[f.pos.row][f.pos.col].Bar(args...)
}

func (f *Figure) Subplot(x, y, pos int) {
	f.rows = x
	f.cols = y

	f.pos = toPositon(x, y, pos)
	f.Align = makeAlign(f.Align, x, y)

	if f.Align[f.pos.row][f.pos.col] == nil {
		f.Align[f.pos.row][f.pos.col] = NewAxis()
	}
}

func makeAlign(align [][]*Axis, x, y int) [][]*Axis {
	rows := len(align)
	if rows != x {
		align = make([][]*Axis, x)
	}
	for j := 0; j < rows; j++ {
		if len(align[j]) != y {
			align[j] = make([]*Axis, y)
		}
	}
	return align
}

func toPositon(x, y, pos int) Positon {
	var col int
	row := pos / y
	if pos%y == 0 {
		col = y - 1
		row = row - 1
	} else {
		col = pos%y - 1
	}
	return Positon{
		row: row,
		col: col,
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
				plots[j][i] = axis.p
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
	w := f.app.NewWindow("Figure")
	w.SetContent(image)
	w.SetPadded(false)
	w.ShowAndRun()
}
