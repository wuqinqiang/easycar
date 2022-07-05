package grpc

import (
	"fmt"
	"strings"
)

var _ Parser = &Default{}

type Default struct {
	uri string
}

func NewDefault(uri string) *Default {
	return &Default{uri: uri}
}

func (d *Default) Get() (service string, method string, err error) {
	sep := strings.IndexByte(d.uri, '/')
	if sep < 0 {
		return "", "", fmt.Errorf("bad url: '%s'. no '/' found", d.uri)
	}
	if strings.Contains(d.uri, "://") {
		return "", "", fmt.Errorf("call dtmdriver.Use() before you use custom scheme for '%s'", d.uri)
	}
	return d.uri[:sep], d.uri[sep:], nil
}
