package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wuqinqiang/easycar/conf"
)

type File struct {
	path string
	conf.EasyCar
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Load() (conf.Loader, error) {
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
	defer file.Close() //nolint:errcheck
	byteAll, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteAll, &f.EasyCar)
	return f, err
}
