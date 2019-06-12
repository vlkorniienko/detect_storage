package detect_storage

import (
	"os/user"
	"io/ioutil"
	"strings"
)

func DetectRemovableStorage() ([]string, error) {
	var drives []string
	var directories []string

	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	pathToSrorageDir := "/media/" + strings.ToLower(user.Name)
	fileInfo, err := ioutil.ReadDir(pathToSrorageDir)
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		directories = append(directories, file.Name())
	}
	for i := range directories {
			drives = append(drives, "/media/" + strings.ToLower(user.Name) + "/" + directories[i])
		}
	return drives, nil
}