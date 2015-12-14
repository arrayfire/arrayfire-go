package arrayfire

/*
#include <arrayfire.h>
*/
import "C"
import (
	"errors"
)

var (
	ErrPrint = errors.New("Failed: af_print_array_gen()")
)

func Print(a Array) error {
	aferr := C.af_print_array_gen(C.CString(""), (C.af_array)(a.arr), 4)
	if aferr != 0 {
		return ErrGetDeviceCount
	}
	return nil
}
