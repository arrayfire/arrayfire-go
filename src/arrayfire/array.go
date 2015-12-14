package arrayfire

/*
#include <arrayfire.h>
#include <af/array.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type (
	DType C.af_dtype
	Dim   C.dim_t
)

type array struct {
	arr C.af_array
}

type Array *array

func release(a Array) (e error) {
	e = nil
	if a.arr != nil {
		e = af_call(C.af_release_array((C.af_array)(a.arr)))
		a.arr = nil
	}
	return
}

func register(in array) (out Array) {

	//TODO: Call runtime.GC() depending on how much memory is left on device

	out = &in
	runtime.SetFinalizer(out, release)
	return
}

func CreateArray(data unsafe.Pointer, ndims uint, dims []Dim, ty DType) (out Array, err error) {

	out = nil
	err = nil

	var a array
	err = af_call(C.af_create_array(&a.arr, data, (C.uint)(ndims),
		(*C.dim_t)(&dims[0]), (C.af_dtype)(ty)))

	out = register(a)
	return
}

func CopyArray(in Array) (out Array, err error) {

	out = nil
	err = nil

	var a array
	err = af_call(C.af_copy_array(&a.arr, in.arr))

	out = register(a)
	return
}

func RetainArray(in Array) (out Array, err error) {

	out = nil
	err = nil

	var a array
	err = af_call(C.af_retain_array(&a.arr, in.arr))

	out = register(a)
	return
}

/*

af_err af_create_handle(af_array *arr, const unsigned ndims, const dim_t * const dims, const af_dtype type)
{
    return CALL(arr, ndims, dims, type);
}

af_err af_write_array(af_array arr, const void *data, const size_t bytes, af_source src)
{
    return CALL(arr, data, bytes, src);
}

af_err af_get_data_ptr(void *data, const af_array arr)
{
    return CALL(data, arr);
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
