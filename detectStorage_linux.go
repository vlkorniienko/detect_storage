package detect_storage

import (
	"os/user"
	"io/ioutil"
	"strings"
	"path"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string

	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	pathToSrorageDir := path.Join("/media/", strings.ToLower(user.Name))
	fileInfo, err := ioutil.ReadDir(pathToSrorageDir)
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		storagePath := path.Join(pathToSrorageDir, file.Name())
		drives = append(drives, storagePath)
	}
	return drives, nil
}