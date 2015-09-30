package arrayfire

/*
#include <arrayfire.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrGetDeviceCount = errors.New("Failed: af_get_device_count()")
	ErrGetDeviceInfo  = errors.New("Failed: af_get_device_info()")
	ErrSetDevice      = errors.New("Failed: af_set_device()")
)

type DeviceInfo struct {
	DName, DPlatform, Toolkit, Compute string
}

func SetDevice(num int) error {
	aferr := C.af_set_device((C.int)(num))
	if aferr != 0 {
		return ErrSetDevice
	}
	return nil
}

// Info returns the name, platform, toolkit, and compute identifiers
func GetDeviceInfo() (DeviceInfo, error) {

	var info DeviceInfo
	cdname := C.CString(info.DName)
	cdplatform := C.CString(info.DPlatform)
	ctoolkit := C.CString(info.Toolkit)
	ccompute := C.CString(info.Compute)

	defer func() {
		C.free(unsafe.Pointer(cdname))
		C.free(unsafe.Pointer(cdplatform))
		C.free(unsafe.Pointer(ctoolkit))
		C.free(unsafe.Pointer(ccompute))
	}()

	aferr := C.af_device_info(cdname, cdplatform, ctoolkit, ccompute)

	if aferr != 0 {
		return info, ErrGetDeviceInfo
	}

	info.DName = C.GoString(cdname)
	info.DPlatform = C.GoString(cdplatform)
	info.Toolkit = C.GoString(ctoolkit)
	info.Compute = C.GoString(ccompute)

	return info, nil
}

// GetDeviceCount returned the device count
func GetDeviceCount() (int, error) {
	var cnt int
	aferr := C.af_get_device_count((*C.int)(unsafe.Pointer(&cnt)))
	if aferr != 0 {
		return 0, ErrGetDeviceCount
	}
	return cnt, nil
}
