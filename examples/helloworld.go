package main

import (
	af "arrayfire"
	"fmt"
	"unsafe"
)

func main() {

	af.Info()

	var count int
	var err error

	count, err = af.GetDeviceCount()
	if err != nil {
		panic(fmt.Sprintf("%s\n", err))
	}
	fmt.Printf("Number of devices: %d\n", count)

	for idx := 0; idx < count; idx++ {
		var info af.DeviceInfo
		af.SetDevice(idx)
		info, err = af.GetDeviceInfo()

		if err != nil {
			panic(fmt.Sprintf("failed after info.Info %s:\n", err))
		}

		fmt.Printf("%d. Device: %s, Platform: %s, Toolkit: %s, Compute: %s\n",
			idx, info.DName, info.DPlatform, info.Toolkit, info.Compute)
	}

	af.SetDevice(0)

	var ndims uint = 1
	var atyp af.DType = af.U32
	var data [][]uint32 = [][]uint32{
		{1, 2, 3, 4, 5},
	}

	var a af.Array
	var dims []af.Dim = []af.Dim{5}
	a, err = af.CreateArray((unsafe.Pointer)(&data[0][0]), ndims, dims, atyp)

	if err != nil {
		panic(fmt.Sprintf("failed at s:\n", err))
	}

	err = af.Print(a)
	if err != nil {
		panic(fmt.Sprintf("failed at %s:\n", err))
	}
}
