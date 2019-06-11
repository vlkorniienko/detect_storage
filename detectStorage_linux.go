package detect_storage

import (
	"os/user"
	"os/exec"
	"strings"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string

	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	lsExecPath := "ls /media/" + strings.ToLower(user.Name)
	lsMediadir := exec.Command("sh", "-c", lsExecPath)
	stdoutMediaDir, err := lsMediadir.CombinedOutput()
	if err != nil {
		return nil, err
	}
	mediaDirArray := strings.Split(string(stdoutMediaDir), "\n")
	for i := range mediaDirArray {
		if mediaDirArray[i] != "" {
			storageDirpath := "/media/" + strings.ToLower(user.Name) + "/" + mediaDirArray[i]
			drives = append(drives, storageDirpath)
		}
	}
	return drives, nil
}