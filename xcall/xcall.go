package xcall

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Foundation
#import <Foundation/Foundation.h>
#include <stdlib.h>
extern int main(int argc, char *argv[]);
*/
import "C"
import (
	"strings"
	"unsafe"
)

// Call the obj-C xcall code to open the passed URI as an x-callback-url
func Call(uri string) {
	cstr := C.CString(strings.Join(
		[]string{"-url bear://x-callback-url", uri}, "/",
	))
	C.main(0, &cstr)
	defer C.free(unsafe.Pointer(cstr))
}
