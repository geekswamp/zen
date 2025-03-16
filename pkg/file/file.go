package file

import "os"

var StatFunc = os.Stat

func IsExist(path string) (os.FileInfo, bool) {
	f, err := StatFunc(path)
	return f, err == nil || os.IsExist(err)
}
