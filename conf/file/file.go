package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wuqinqiang/easycar/conf"
)

var _ conf.Source = &File{}

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Load() (*conf.KeyValue, error) {
	fi, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("the path must be a file,it's a dir now")
	}
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteAll, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &conf.KeyValue{
		Value:  byteAll,
		Format: GetFormatByFileName(fileInfo.Name()),
	}, nil
}

func GetFormatByFileName(fileName string) string {
	spilt := strings.Split(fileName, ".")
	if len(spilt) > 1 {
		return spilt[len(spilt)-1]
	}
	return ""
}
