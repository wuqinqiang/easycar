package protocol

type Resp struct {
	Code int64  //http code
	Body []byte // response body
}