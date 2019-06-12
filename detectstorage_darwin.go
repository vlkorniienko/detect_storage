package detect_storage

import (
	"io/ioutil"
	"path"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string

	const pathToStoragesDir = "/Volumes"
	fileInfo, err := ioutil.ReadDir(pathToStoragesDir)
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		if file.Name() != "Macintosh HD" {
			storagePath := path.Join("/Volumes", file.Name())
			drives = append(drives, storagePath)
		}
	}
	return drives, nil
}
