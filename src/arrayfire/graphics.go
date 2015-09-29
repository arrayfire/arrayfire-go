package arrayfire

/*
#include <arrayfire.h>
#include <af/graphics.h>
extern AFAPI af_err 	af_device_info (char *d_name, char *d_platform, char *d_toolkit, char *d_compute);
extern AFAPI int 	getDeviceCount();
*/
// #cgo LDFLAGS: -L/usr/local/lib -lafcuda
import "C"
import "errors"

type (
	AFWindow C.af_window
	AFCell   C.af_cell
)

type Window struct {
	window C.af_window
	width  int
	height int
	title  string
}

var (
	ErrNone           = errors.New("No Error")
	ErrCreateWindow   = errors.New("Failed: af_create_window()")
	ErrSetPosition    = errors.New("Failed: af_set_position()")
	ErrSetTitle       = errors.New("Failed: af_set_title()")
	ErrGrid           = errors.New("Failed: af_grid()")
	ErrShow           = errors.New("Failed: af_show()")
	ErrDrawPlot       = errors.New("Failed: af_draw_plot()")
	ErrSetSize        = errors.New("Failed: af_set_size()")
	ErrIsWindowClosed = errors.New("Failed: af_is_window_closed()")
	ErrDrawImage      = errors.New("Failed: af_draw_image()")
	ErrDrawHist       = errors.New("Failed: af_draw_hist()")
)

func (w *Window) Create(width, height int, title string) error {
	w.width = width
	w.height = height
	w.title = title

	aferr := C.af_create_window(&w.window, C.int(w.width), C.int(w.height), C.CString(w.title))

	if aferr != 0 {
		return ErrCreateWindow
	}

	return nil
}

func (w *Window) SetPosition(x, y uint) error {
	aferr := C.af_set_position(w.window, C.uint(x), C.uint(y))
	if aferr != 0 {
		return ErrSetPosition
	}
	return nil
}

func (w *Window) SetTitle(title string) error {
	aferr := C.af_set_title(w.window, C.CString(title))
	if aferr != 0 {
		return ErrSetTitle
	}

	return nil
}

func (w *Window) Grid(rows, cols int) error {
	aferr := C.af_grid(w.window, C.int(rows), C.int(cols))

	if aferr != 0 {
		return ErrGrid
	}

	return nil
}

func (w *Window) Show() error {
	aferr := C.af_show(w.window)
	if aferr != 0 {
		return ErrShow
	}

	return nil
}

func (w *Window) DrawPlot(x Array, y Array, props *AFCell) error {
	aferr := C.af_draw_plot(w.window, (C.af_array)(x._array), (C.af_array)(y._array), (*C.af_cell)(props))
	if aferr != 0 {
		return ErrDrawPlot
	}

	return nil
}

func (w *Window) SetSize(width, height uint) error {
	aferr := C.af_set_size(w.window, C.uint(width), C.uint(height))
	if aferr != 0 {
		return ErrSetSize
	}
	return nil
}

func (w *Window) IsClosed() (bool, error) {
	var closed bool
	aferr := C.af_is_window_closed((*C._Bool)(&closed), w.window)
	if aferr != 0 {
		return false, ErrIsWindowClosed
	}

	return closed, nil
}

func (w *Window) DrawImage(a Array, props AFCell) error {
	aferr := C.af_draw_image(w.window, (C.af_array)(a._array), (*C.af_cell)(&props))
	if aferr != 0 {
		return ErrDrawImage
	}
	return nil
}

func (w *Window) DrawHist(x Array, minval, maxval float64, props AFCell) error {
	aferr := C.af_draw_hist(w.window, (C.af_array)(x._array), C.double(minval), C.double(maxval), (*C.af_cell)(&props))
	if aferr != 0 {
		return ErrDrawHist
	}
	return nil
}
