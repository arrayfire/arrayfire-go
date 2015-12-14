package arrayfire

/*
#include <arrayfire.h>
#include <af/array.h>
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

var (
	ErrAfCreateArray = errors.New("Failed: af_create_array()")
	ErrRelease       = errors.New("Failed: af_release_array()")
)

type (
	DType C.af_dtype
	Dim   C.dim_t
)

type array struct {
	arr C.af_array
}

type Array *array

func release(a Array) error {
	if a.arr != nil {
		aferr := C.af_release_array((C.af_array)(a.arr))
		if aferr != 0 {
			return ErrRelease
		}
		a.arr = nil
	}
	return nil
}

func register(a array) (r Array) {

	//TODO: Call runtime.GC() depending on how much memory is left on device

	r = &a
	runtime.SetFinalizer(r, release)
	return
}

func CreateArray(data unsafe.Pointer, ndims uint, dims []Dim, ty DType) (r Array, err error) {

	r = nil
	err = nil

	var a array
	aferr := C.af_create_array(&a.arr, data, (C.uint)(ndims), (*C.dim_t)(&dims[0]), (C.af_dtype)(ty))

	if aferr != 0 {
		err = ErrAfCreateArray
	}

	r = register(a)
	return
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
