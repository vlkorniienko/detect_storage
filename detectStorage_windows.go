package detect_storage

import (
	"unsafe"
	"golang.org/x/sys/windows"
)

const DRIVE_REMOVABLE = 2

func DetectRemovableStorage() ([]string, error) {
	kernelDll := windows.NewLazyDLL("kernel32.dll")
	pGetDriveTypeA := kernelDll.NewProc("GetDriveTypeA")
	pGetLogicalDrives := kernelDll.NewProc("GetLogicalDrives")
	var drives []string

	bitMap, _, err := pGetLogicalDrives.Call()
	if err == nil {
		return nil, err
	} else {
		for i := 'A'; i <= 'Z'; i++ {
			if bitMap&1 == 1 {
				droot := string(i) + ":\\"
				tp, _, _ := pGetDriveTypeA.Call(uintptr(unsafe.Pointer(&[]byte(droot)[0])))
				if tp == DRIVE_REMOVABLE {
					drives = append(drives, string(i) + ":\\")
				}
			}
			bitMap >>= 1
		}
	}
	return drives, nil
}
