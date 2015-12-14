package main

import (
	af "arrayfire"
	"fmt"
	"unsafe"
)

func main() {

	af.Info()

	var err error

	var ndims uint = 1
	var data [][]uint32 = [][]uint32{
		{1, 2, 3, 4, 5},
	}

	var a af.Array
	var dims []af.Dim = []af.Dim{5}
	a, err = af.CreateArray((unsafe.Pointer)(&data[0][0]), ndims, dims, af.U32)

	if err != nil {
		panic(fmt.Sprintf("failed at s:\n", err))
	}

	err = af.Print(a)
	if err != nil {
		panic(fmt.Sprintf("failed at %s:\n", err))
	}
}
