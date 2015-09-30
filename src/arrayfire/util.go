package arrayfire

/*
#include <arrayfire.h>
*/
import "C"
import (
	"errors"
)

var (
	ErrPrint = errors.New("Failed: af_print_array()")
	ErrInfo  = errors.New("Failed: af_get_device_info()")
)

func Info() error {
	aferr := C.af_info()
	if aferr != 0 {
		return ErrInfo
	}
	return nil
}

func Print(arr Array) error {
	aferr := C.af_print_array((C.af_array)(arr))
	if aferr != 0 {
		return ErrGetDeviceCount
	}
	return nil
}
