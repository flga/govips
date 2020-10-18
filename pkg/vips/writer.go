package vips

// #cgo pkg-config: vips
// #include "bridge.h"
import "C"
import (
	"fmt"
	"io"
	"unsafe"

	"github.com/mattn/go-pointer"
)

//export govips_obs_write
func govips_obs_write(target *C.VipsTargetCustom, buffer unsafe.Pointer, length C.long, user unsafe.Pointer) C.long {
	v := pointer.Restore(user).(io.Writer)
	fmt.Println("write called with", length)
	d := C.GoBytes(buffer, C.int(length))
	n, err := v.Write(d)
	if err != nil {
		return -1
	}
	return C.long(n)
}

//export govips_obs_finish
func govips_obs_finish(target *C.VipsTargetCustom, user unsafe.Pointer) {
	fmt.Println("finish called")
	u := (interface{})(user)
	if flusher, ok := u.(interface{ Flush() }); ok {
		flusher.Flush()
	}
	if closer, ok := u.(io.Closer); ok {
		closer.Close()
	}
}
