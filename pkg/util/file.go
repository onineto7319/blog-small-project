package util

import (
	"io/ioutil"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetFilesName(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	filesName := make([]string, 0)
	for _, f := range files {
		filesName = append(filesName, f.Name())
	}
	return filesName, nil
}

func GetCurrentFileDir() string {
	dir, _ := os.Getwd()
	return dir
}
