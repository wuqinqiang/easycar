package file

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/wuqinqiang/easycar/conf/common"
)

type File struct {
	path string
	*common.EasyCar
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Load() (*common.EasyCar, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(dir + f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close() //nolint:errcheck
	byteAll, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(byteAll, &f.EasyCar)
	return f.EasyCar, err
}
