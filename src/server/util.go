package server

import (
	"io/fs"
)

func FilterFileList(slice []fs.FileInfo) []string {
	var s []string
	for _, item := range slice {
		if item.IsDir() == false {
			s = append(s, item.Name())
		}
	}

	return s
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
