package main

import (
	"arrayfire"
	"fmt"
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
}
