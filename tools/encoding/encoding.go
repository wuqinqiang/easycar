package encoding

import "strings"

type Codec interface {
	Name() string
	Marshal(res interface{}) ([]byte, error)
	UnMarshal(data []byte, src interface{}) error
}

var (
	codecMap = make(map[string]Codec)
)

func RegisterCodec(codec Codec) {
	if codec == nil {
		return
	}
	if codec.Name() == "" {
		return
	}
	name := strings.ToLower(codec.Name())
	codecMap[name] = codec
}

func GetCodecByFormat(name string) (codec Codec, ok bool) {
	codec, ok = codecMap[name]
	return
}
