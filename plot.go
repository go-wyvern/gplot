package gplot

import (
	"log"
	"reflect"
	"unsafe"

	"gonum.org/v1/gonum/floats"
)

const (
	GopPackage = true
	Gop_game   = "Figure"
	Gop_sprite = "Axis"
)

type Ploter interface {
	initPlot()
	finishPlot()
}

// Gopt_Figure_Main is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Main(plot Ploter) {
	plot.initPlot()
	defer plot.finishPlot()
	plot.(interface{ MainEntry() }).MainEntry()
}

// Gopt_Figure_Run is required by Go+ compiler as the entry of a .plot project.
func Gopt_Figure_Run(plot Ploter, x, y int) {
	v := reflect.ValueOf(plot).Elem()
	t := reflect.TypeOf(plot).Elem()
	p := instance(v)
	for i, n := 0, v.NumField(); i < n; i++ {
		typ := t.Field(i).Type
		m, ok := reflect.PtrTo(typ).MethodByName("Main")
		if ok {
			axis := reflect.ValueOf(NewAxis("title", "x", "y"))
			axis2 := reflect.New(typ)
			axis3 := axis2.Elem().FieldByName("Axis")
			axis3.Set(axis.Elem())
			ax := axis3.Addr().Interface().(*Axis)
			m.Func.Call([]reflect.Value{axis2})
			p.Subplot(x, y, ax.pos)
			p.Align[p.pos.row][p.pos.col] = ax
		}
	}
}

// Plot gop enter func
func Plot__0(plot Ploter, args ...[]float64) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Plot__0(args...)
}

func Plot__1(plot Ploter, args ...[]int) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Plot__1(args...)
}

// Plot gop enter func
func Plot__2(plot Ploter, args ...Vector) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Plot__2(args...)
}

func Bar(plot Ploter, args ...interface{}) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Bar(args...)
}

func Subplot(plot Ploter, row, col, pos int) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Subplot(row, col, pos)
}

func Legend(plot Ploter, labels ...string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Legend(labels...)
}

func Xlabel(plot Ploter, xlabel string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Xlabel(xlabel)
}

func Ylabel(plot Ploter, ylabel string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Ylabel(ylabel)
}

func Title(plot Ploter, title string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.Title(title)
}

func NominalX(plot Ploter, names ...string) {
	v := reflect.ValueOf(plot).Elem()
	p := instance(v)
	p.NominalX(names...)
}

func Linspace(l, r float64, n int) Vector {
	s := make([]float64, n)
	dst := floats.Span(s, l, r)
	return newVec(dst)
}

func instance(plotter reflect.Value) *Figure {
	fld := plotter.FieldByName("Figure")
	if !fld.IsValid() {
		log.Panicf("type %v doesn't has field gplot.Figure", plotter.Type())
	}
	return fld.Addr().Interface().(*Figure)
}

func getFieldPtrOrAlloc(v reflect.Value, i int) (name string, val interface{}) {
	tFld := v.Type().Field(i)
	vFld := v.Field(i)
	typ := tFld.Type
	word := unsafe.Pointer(vFld.Addr().Pointer())
	ret := makeEmptyInterface(reflect.PtrTo(typ), word)
	return tFld.Name, ret
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func makeEmptyInterface(typ reflect.Type, word unsafe.Pointer) (i interface{}) {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	etyp := (*emptyInterface)(unsafe.Pointer(&typ))
	e.typ, e.word = etyp.word, word
	return
}
