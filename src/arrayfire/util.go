package arrayfire

/*
#include <arrayfire.h>
*/
import "C"

func Print(a Array) error {
	return af_call(C.af_print_array_gen(C.CString(""), (C.af_array)(a.arr), 4))
}
