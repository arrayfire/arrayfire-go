package main

import (
	af "arrayfire"
	"fmt"
	"time"
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
		fmt.Printf("%d. Device: %s, Platform: %s, Toolkit: %s, Compute: %s\n", idx, info.DName, info.DPlatform, info.Toolkit, info.Compute)
	}

	af.SetDevice(0)

	var ndims uint = 1
	var atyp af.DType = af.U32
	var data [][]uint32 = [][]uint32{
		{1, 2, 3, 4, 5},
	}

	var a af.Array
	var dims []af.Dim = []af.Dim{5}
	err = af.Create(&a, (unsafe.Pointer)(&data[0][0]), ndims, dims, atyp)

	if err != nil {
		panic(fmt.Sprintf("failed at s:\n", err))
	}

	err = af.Print(a)
	if err != nil {
		panic(fmt.Sprintf("failed at %s:\n", err))
	}

	defer func() {
		err = af.Release(a)
		if err != nil {
			fmt.Printf("failed at %s:\n", err)
		}
	}()

	if false {
		var w af.Window
		err = w.Create(400, 400, "test!")
		if err != nil {
			fmt.Printf("failed %s:\n", err)
		} else {
			fmt.Printf("Created window\n")
		}

		err = w.SetPosition(400, 400)
		if err != nil {
			fmt.Printf("failed %s:\n", err)
		} else {
			fmt.Printf("Set window position\n")
		}

		err = w.SetTitle("Testing Title")
		if err != nil {
			fmt.Printf("failed %s:\n", err)
		} else {
			fmt.Printf("Set window title\n")
		}

		err = w.Grid(10, 10)
		if err != nil {
			fmt.Printf("failed %s:\n", err)
		} else {
			fmt.Printf("Set window grid\n")
		}

		err = w.Show()
		if err != nil {
			fmt.Printf("failed %s:\n", err)
		} else {
			fmt.Printf("show window\n")
		}
	}

	fmt.Printf("Waiting for 10 seconds\n")
	time.Sleep(10 * time.Second)
}
