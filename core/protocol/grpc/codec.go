package grpc

import "fmt"

type rawCodec struct{}

func (cb rawCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (cb rawCodec) Unmarshal(data []byte, v interface{}) error {
	ba, ok := v.(*[]byte)
	if !ok {
		return fmt.Errorf("please pass in *[]byte")
	}
	*ba = append(*ba, data...)
	return nil
}

func (cb rawCodec) Name() string { return "easycar" }
