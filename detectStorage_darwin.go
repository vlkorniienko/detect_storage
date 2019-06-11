package detect_storage

import (
	"os/exec"
	"strings"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string

	cmd := exec.Command("sh", "-c", "ls /Volumes")
	stdoutVolumeDir, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	volumeDirArray := strings.Split(string(stdoutVolumeDir), "\n")
	for i := range volumeDirArray {
		if volumeDirArray[i] != "Macintosh HD" && volumeDirArray[i] != "" {
			storagePath := "/Volumes/" + volumeDirArray[i]
			drives = append(drives, storagePath)
		}
	}
	return drives, nil
}

