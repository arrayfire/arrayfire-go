package main

import (
	"arrayfire"
	"fmt"
	"time"
)

func main() {

	var af arrayfire.AFInfo
	var a arrayfire.AFArray

	fmt.Printf("a is of type %T\n", a)
	err := af.Info()

	if err != nil {
		fmt.Printf("failed %s:\n", err)
	} else {
		fmt.Printf("%s, %s, %s, %s\n", af.DName, af.DPlatform, af.Toolkit, af.Compute)
	}

	count, err := af.GetDeviceCount()
	if err != nil {
		fmt.Printf("failed %s:\n", err)
	} else {
		fmt.Printf("device count %d=%d\n", count, af.Count)
	}

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

	time.Sleep(10 * time.Second)
}
