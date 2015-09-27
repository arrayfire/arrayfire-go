package arrayfire

/*
#include <arrayfire.h>
#include <af/util.h>
extern AFAPI af_err 	af_device_info (char *d_name, char *d_platform, char *d_toolkit, char *d_compute);
extern AFAPI int 	getDeviceCount();
*/
// #cgo LDFLAGS: -L/usr/local/lib -lafcuda
import "C"
import (
	"errors"
	"unsafe"
)

var (
	// ErrGetDeviceCount is returned when the af_get_device_count() fails
	ErrGetDeviceCount = errors.New("Failed: af_get_device_count()")
	// ErrGetDeviceInfo is returned when the af_get_device_info() fails
	ErrGetDeviceInfo = errors.New("Failed: af_get_device_info()")
)

// AFInfo contains info from call
type AFInfo struct {
	DName, DPlatform, Toolkit, Compute string
	Count                              int
}

// Info returns the name, platform, toolkit, and compute identifiers
func (af *AFInfo) Info() error {

	cdname := C.CString(af.DName)
	cdplatform := C.CString(af.DPlatform)
	ctoolkit := C.CString(af.Toolkit)
	ccompute := C.CString(af.Compute)

	defer func() {
		C.free(unsafe.Pointer(cdname))
		C.free(unsafe.Pointer(cdplatform))
		C.free(unsafe.Pointer(ctoolkit))
		C.free(unsafe.Pointer(ccompute))
	}()

	aferr := C.af_device_info(cdname, cdplatform, ctoolkit, ccompute)

	if aferr != 0 {
		return ErrGetDeviceInfo
	}

	af.DName = C.GoString(cdname)
	af.DPlatform = C.GoString(cdplatform)
	af.Toolkit = C.GoString(ctoolkit)
	af.Compute = C.GoString(ccompute)

	// fmt.Printf("[%s, %s, %s, %s]\n", af.DName, af.DPlatform, af.Toolkit, af.Compute)

	return nil
}

// GetDeviceCount returned the device count
func (af *AFInfo) GetDeviceCount() (int, error) {
	var cnt int
	aferr := C.af_get_device_count((*C.int)(unsafe.Pointer(&cnt)))
	if aferr != 0 {
		return 0, ErrGetDeviceCount
	}
	af.Count = cnt
	return cnt, nil
}
