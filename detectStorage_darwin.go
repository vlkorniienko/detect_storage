package detect_storage

import (
	"io/ioutil"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string
	var directories []string

	const pathToStoragesDir = "/Volumes"
	fileInfo, err := ioutil.ReadDir(pathToStoragesDir)
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		directories = append(directories, file.Name())
	}
	for i := range directories {
		if directories[i] != "Macintosh HD" {
			drives = append(drives, "/Volumes/" + directories[i])
		}
	}
	return drives, nil
}
