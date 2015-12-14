# ArrayFire Go Bindings

This project is still in the prototype phase. Please follow the discussionon [this issue](https://github.com/arrayfire/arrayfire-go/issues/1). You can also chat with us at the following locations:

[![Join the chat at https://gitter.im/arrayfire/arrayfire-go](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/arrayfire/arrayfire-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Building arrayfire-go

#### Install ArrayFire libraries
Ensure you have the ArrayFire library installed on your system.

You can get the ArrayFire library from one of the following ways:
 - [Binary Installers](http://arrayfire.com/download)
 - [Build from source](http://github.com/arrayfire/arrayfire)

#### Set Environment variables

Point `AF_PATH` to the **installed** location of ArrayFire.

    $ export AF_PATH=/path/to/arrayfire


If using the CUDA backend, you will need to include path to nvvm

    $ export NVVM_LIB_PATH=/path/to/cuda/nvvm/lib64 # use nvvm/lib for OSX or 32 bit systems


Point `GOPATH` to the location of arrayfire-go

Export `LD_LIBRARY_PATH` on Linux and `DYLD_LIBRARY_PATH` on OSX to point to
- `AF_PATH/lib`
- `NVVM_LIB_PATH` if using CUDA backend.

#### Building with arrayfire libraries

For ArrayFire 3.2 and above:

    $ export CGO_CFLAGS="-I$AF_PATH/include"
    $ export CGO_LDFLAGS="-L$AF_PATH/lib -laf -lforge"
    $ go build

For older versions replace `-laf` with either `-lafcuda`, `-lafopencl` or `-lafcpu`

## Contribute

If you want to contribute to this project, start here:
+ https://github.com/golang/go/wiki/cgo
+ https://golang.org/cmd/cgo/
+ http://blog.golang.org/c-go-cgo
