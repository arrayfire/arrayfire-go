package arrayfire

/*
#include <arrayfire.h>
*/
import "C"
import (
	"unsafe"
)

type DeviceInfo struct {
	DName, DPlatform, Toolkit, Compute string
}

func Info() error {
	return af_call(C.af_info())
}

func SetDevice(num int) error {
	return af_call(C.af_set_device((C.int)(num)))
}

// Info returns the name, platform, toolkit, and compute identifiers
func GetDeviceInfo() (info DeviceInfo, e error) {

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

	e = af_call(C.af_device_info(cdname, cdplatform, ctoolkit, ccompute))

	info.DName = C.GoString(cdname)
	info.DPlatform = C.GoString(cdplatform)
	info.Toolkit = C.GoString(ctoolkit)
	info.Compute = C.GoString(ccompute)

	return
}

// GetDeviceCount returned the device count
func GetDeviceCount() (count int, e error) {
	e = af_call(C.af_get_device_count((*C.int)(unsafe.Pointer(&count))))
	return
}
