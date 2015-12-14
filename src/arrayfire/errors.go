package arrayfire

/*
#include <arrayfire.h>
#include <af/array.h>
*/
import "C"
import (
	"errors"
)

var (
	ErrNoMemory      = errors.New("Out of device memory")
	ErrDriver        = errors.New("Driver error")
	ErrRuntime       = errors.New("Runtime error")
	ErrInvalidArray  = errors.New("Input is an invalid Array")
	ErrArgument      = errors.New("Value of input argument is incorrect")
	ErrSize          = errors.New("Size of input argument is incorrect")
	ErrType          = errors.New("Type of input argument is incorrect")
	ErrTypeMismatch  = errors.New("Mismatch between input types")
	ErrBatch         = errors.New("Batch mode / GFOR not supported")
	ErrNotSupported  = errors.New("Option not supported")
	ErrNotConfigured = errors.New("ArrayFire not built with this feature")
	ErrNonFree       = errors.New("Function requires nonfree build of ArrayFire")
	ErrNoDouble      = errors.New("Device does not support double precision")
	ErrNoGraphics    = errors.New("ArrayFire not built with graphics")
	ErrLoadLibrary   = errors.New("Error while loading library")
	ErrLoadSymbol    = errors.New("Error while loading symbol")
	ErrInternal      = errors.New("Internal error")
	ErrUnknown       = errors.New("Unknown error")
)

const (
	SUCCESS            = 0
	ERR_NO_MEM         = 101
	ERR_DRIVER         = 102
	ERR_RUNTIME        = 103
	ERR_INVALID_ARRAY  = 201
	ERR_ARG            = 202
	ERR_SIZE           = 203
	ERR_TYPE           = 204
	ERR_DIFF_TYPE      = 205
	ERR_BATCH          = 207
	ERR_NOT_SUPPORTED  = 301
	ERR_NOT_CONFIGURED = 302
	ERR_NONFREE        = 303
	ERR_NO_DBL         = 401
	ERR_NO_GFX         = 402
	ERR_LOAD_LIB       = 501
	ERR_LOAD_SYM       = 502
	ERR_INTERNAL       = 998
	ERR_UNKNOWN        = 999
)

func af_call(e C.af_err) (err error) {
	switch e {
	case SUCCESS:
		return nil
	case ERR_NO_MEM:
		return ErrNoMemory
	case ERR_DRIVER:
		return ErrDriver
	case ERR_RUNTIME:
		return ErrRuntime
	case ERR_INVALID_ARRAY:
		return ErrInvalidArray
	case ERR_ARG:
		return ErrArgument
	case ERR_SIZE:
		return ErrSize
	case ERR_TYPE:
		return ErrType
	case ERR_DIFF_TYPE:
		return ErrTypeMismatch
	case ERR_BATCH:
		return ErrBatch
	case ERR_NOT_SUPPORTED:
		return ErrNotSupported
	case ERR_NOT_CONFIGURED:
		return ErrNotConfigured
	case ERR_NONFREE:
		return ErrNonFree
	case ERR_NO_DBL:
		return ErrNoDouble
	case ERR_NO_GFX:
		return ErrNoGraphics
	case ERR_LOAD_LIB:
		return ErrLoadLibrary
	case ERR_LOAD_SYM:
		return ErrLoadSymbol
	case ERR_INTERNAL:
		return ErrInternal
	case ERR_UNKNOWN:
		return ErrUnknown
	default:
		return ErrUnknown
	}
}
