package file

import (
	"bytes"
	"os"

	"github.com/geekswamp/zen/pkg/errors"
)

var buffSize = 1 << 20

type ReadLineFromEnd struct {
	file     *os.File
	fileSize int
	buff     *bytes.Buffer
	lineBuff []byte
	swapBuff []byte
	isFirst  bool
}

func NewReadLineFromEnd(filename string) (rd *ReadLineFromEnd, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	i, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if i.IsDir() {
		return nil, errors.ErrNotAFile
	}

	fz := int(i.Size())

	rd = &ReadLineFromEnd{
		file:     f,
		fileSize: fz,
		buff:     bytes.NewBuffer([]byte{}),
		lineBuff: make([]byte, 0),
		swapBuff: make([]byte, buffSize),
		isFirst:  true,
	}

	return rd, nil
}

func IsExist(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}
