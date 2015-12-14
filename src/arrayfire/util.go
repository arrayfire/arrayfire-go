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

func Print(arr Array) error {
	aferr := C.af_print_array_gen(C.CString(""), (C.af_array)(arr), 4)
	if aferr != 0 {
		return ErrGetDeviceCount
	}
	return nil
}
