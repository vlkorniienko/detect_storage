package detect_storage

// #include <stdlib.h>
import "C"
import (
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

func DetectRemovableStorage() ([]string, error) {
	const DRIVE_REMOVABLE = 2
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")
	kernelDll := windows.NewLazyDLL("kernel32.dll")
	pGetDriveTypeA := kernelDll.NewProc("GetDriveTypeA")
	var drives []string

	bitMap, _, err := syscall.Syscall(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0)
	if err != 0 {
		return nil, err
	} else {
		for i := 'A'; i <= 'Z'; i++ {
			if bitMap&1 == 1 {
				droot := string(i) + ":\\"
				drootC := C.CString(droot)
				tp, _, _ := pGetDriveTypeA.Call(uintptr(unsafe.Pointer(drootC)))
				defer C.free(unsafe.Pointer(drootC))
				if tp == DRIVE_REMOVABLE {
					drives = append(drives, string(i)+":\\")
				}
			}
			bitMap >>= 1
		}
	}
	return drives, nil
}
