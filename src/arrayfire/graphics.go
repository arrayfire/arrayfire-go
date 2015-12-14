package arrayfire

/*
#include <arrayfire.h>
#include <af/graphics.h>
*/
import "C"

type (
	Cell C.af_cell
)

type Window struct {
	window C.af_window
	width  int
	height int
	title  string
}

func (w *Window) Create(width, height int, title string) error {
	w.width = width
	w.height = height
	w.title = title

	return af_call(C.af_create_window(&w.window, C.int(w.width), C.int(w.height), C.CString(w.title)))
}

func (w *Window) SetPosition(x, y uint) error {
	return af_call(C.af_set_position(w.window, C.uint(x), C.uint(y)))
}

func (w *Window) SetTitle(title string) error {
	return af_call(C.af_set_title(w.window, C.CString(title)))
}

func (w *Window) Grid(rows, cols int) error {
	return af_call(C.af_grid(w.window, C.int(rows), C.int(cols)))
}

func (w *Window) Show() error {
	return af_call(C.af_show(w.window))
}

func (w *Window) DrawPlot(x Array, y Array, props *Cell) error {
	return af_call(C.af_draw_plot(w.window, (C.af_array)(x.arr),
		(C.af_array)(y.arr), (*C.af_cell)(props)))
}

func (w *Window) SetSize(width, height uint) error {
	return af_call(C.af_set_size(w.window, C.uint(width), C.uint(height)))
}

func (w *Window) IsClosed() (closed bool, e error) {
	e = af_call(C.af_is_window_closed((*C._Bool)(&closed), w.window))
	return
}

func (w *Window) DrawImage(a Array, props Cell) error {
	return af_call(C.af_draw_image(w.window, (C.af_array)(a.arr), (*C.af_cell)(&props)))
}

func (w *Window) DrawHist(x Array, minval, maxval float64, props Cell) error {
	return af_call(C.af_draw_hist(w.window, (C.af_array)(x.arr),
		C.double(minval), C.double(maxval), (*C.af_cell)(&props)))
}
