package encoding

var _ Codec = (*jsonCodec)(nil)

type jsonCodec struct {
}

func (j jsonCodec) Name() string {
	return "json"
}

func (j jsonCodec) Marshal(res interface{}) ([]byte, error) {
	panic("implement me")
}

func (j jsonCodec) UnMarshal(data []byte, src interface{}) error {
	panic("implement me")
}

func init() {
	RegisterCodec(jsonCodec{})
}
