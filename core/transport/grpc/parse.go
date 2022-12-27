package grpc

import (
	"fmt"
	"strings"
)

type Parser interface {
	Get() (service string, method string, err error)
}

type parser struct {
	uri string
}

func NewDefault(uri string) Parser {
	return &parser{uri: uri}
}

func (d *parser) Get() (service string, method string, err error) {
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
