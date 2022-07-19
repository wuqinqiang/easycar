package encoding

import "encoding/json"

var _ Codec = (*jsonCodec)(nil)

type jsonCodec struct {
}

func (j jsonCodec) Name() string {
	return "json"
}

func (j jsonCodec) Marshal(res interface{}) ([]byte, error) {
	return json.Marshal(res)
}

func (j jsonCodec) UnMarshal(data []byte, src interface{}) error {
	return json.Unmarshal(data, src)
}

func init() {
	RegisterCodec(jsonCodec{})
}
