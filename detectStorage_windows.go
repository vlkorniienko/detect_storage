package detect_storage

//#include <windows.h>
import "C"
import (
    "syscall"
	"unsafe"
	"strings"
)

func DetectRemovableStorage (drives []string, error) {
    kernel32, _ := syscall.LoadLibrary("kernel32.dll")
    getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    var drives []string

    bitMap, _, callErr := syscall.Syscall(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0)
	if callErr != 0 {
        return nil, callErr
    } else {
         for i := range availableDrives {
			if bitMap & 1 == 1 {
				droot := availableDrives[i] + ":\\"
				drootC := C.CString(droot)
				tp := C.GetDriveTypeA(drootC)
				defer C.free(unsafe.Pointer(drootC))
				if tp == 2 {
					availableDrives[i] += ":\\"
					drives = append(drives, availableDrives[i])
				}
			}
			bitMap >>= 1
		}
		return drives, nil
	}
}
