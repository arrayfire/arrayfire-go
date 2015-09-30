package main

import (
	"arrayfire"
	"fmt"
	"time"
	"unsafe"
)

func main() {

	var af arrayfire.AFInfo
	var a arrayfire.Array

	fmt.Printf("a is of type %T\n", a)
	err := af.Info()

	if err != nil {
		fmt.Printf("failed after af.Info %s:\n", err)
	} else {
		fmt.Printf("%s, %s, %s, %s\n", af.DName, af.DPlatform, af.Toolkit, af.Compute)
	}

	var ndims uint = 1
	var atyp arrayfire.AFDType = arrayfire.U32
	var data [][]uint32 = [][]uint32{
		{1, 2, 3, 4, 5},
	}

	var dims []arrayfire.AFDim = []arrayfire.AFDim{5}
	err = a.Create((unsafe.Pointer)(&data[0][0]), ndims, dims, atyp)
	if err != nil {
		fmt.Printf("failed after create %s:\n", err)
	} else {
		fmt.Printf("array created\n")
	}

	count, err := af.GetDeviceCount()
	if err != nil {
		fmt.Printf("failed %s:\n", err)
	} else {
		fmt.Printf("device count %d=%d\n", count, af.Count)
	}

	if false {
		var w arrayfire.Window
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
