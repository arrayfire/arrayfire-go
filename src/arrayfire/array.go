package arrayfire

/*
#include <arrayfire.h>
#include <af/util.h>
#include <af/array.h>
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
	ErrAfCreateArray = errors.New("Failed: af_create_array()")
)

type (
	_AFArray C.af_array
	_AFDType C.af_dtype
	_AFDim   C.dim_t
)

type AFArray struct {
	_array _AFArray
}

func (ar *AFArray) Create(data unsafe.Pointer, ndims uint, dims *_AFDim, afType _AFDType) error {

	aferr := C.af_create_array((*C.af_array)(ar._array), data, (C.uint)(ndims), (*C.dim_t)(dims), (C.af_dtype)(afType))
	if aferr != 0 {
		return ErrAfCreateArray
	}
	return nil
}

/*

af_err af_create_handle(af_array *arr, const unsigned ndims, const dim_t * const dims, const af_dtype type)
{
    return CALL(arr, ndims, dims, type);
}

af_err af_copy_array(af_array *arr, const af_array in)
{
    return CALL(arr, in);
}

af_err af_write_array(af_array arr, const void *data, const size_t bytes, af_source src)
{
    return CALL(arr, data, bytes, src);
}

af_err af_get_data_ptr(void *data, const af_array arr)
{
    return CALL(data, arr);
}

af_err af_release_array(af_array arr)
{
    return CALL(arr);
}

af_err af_retain_array(af_array *out, const af_array in)
{
    return CALL(out, in);
}

af_err af_get_data_ref_count(int *use_count, const af_array in)
{
    return CALL(use_count, in);
}

af_err af_eval(af_array in)
{
    return CALL(in);
}

af_err af_get_elements(dim_t *elems, const af_array arr)
{
    return CALL(elems, arr);
}

af_err af_get_type(af_dtype *type, const af_array arr)
{
    return CALL(type, arr);
}

af_err af_get_dims(dim_t *d0, dim_t *d1, dim_t *d2, dim_t *d3, const af_array arr)
{
    return CALL(d0, d1, d2, d3, arr);
}

af_err af_get_numdims(unsigned *result, const af_array arr)
{
    return CALL(result, arr);
}

#define ARRAY_HAPI_DEF(af_func) \
af_err af_func(bool *result, const af_array arr)\
{\
    return CALL(result, arr);\
}

ARRAY_HAPI_DEF(af_is_empty)
ARRAY_HAPI_DEF(af_is_scalar)
ARRAY_HAPI_DEF(af_is_row)
ARRAY_HAPI_DEF(af_is_column)
ARRAY_HAPI_DEF(af_is_vector)
ARRAY_HAPI_DEF(af_is_complex)
ARRAY_HAPI_DEF(af_is_real)
ARRAY_HAPI_DEF(af_is_double)
ARRAY_HAPI_DEF(af_is_single)
ARRAY_HAPI_DEF(af_is_realfloating)
ARRAY_HAPI_DEF(af_is_floating)
ARRAY_HAPI_DEF(af_is_integer)
ARRAY_HAPI_DEF(af_is_bool)
*/
