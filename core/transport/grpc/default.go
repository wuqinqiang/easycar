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
	uri := d.uri
	if strings.HasPrefix(d.uri, "grpc://") {
		uri = uri[7:]
	}
	sep := strings.IndexByte(uri, '/')
	if sep < 0 {
		return "", "", fmt.Errorf("bad url: '%s'. no '/' found", d.uri)
	}
	return uri[:sep], uri[sep:], nil
}
