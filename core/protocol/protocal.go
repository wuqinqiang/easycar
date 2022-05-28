package protocol

type Protocol interface {
	Req() (Resp, error)
}

type Resp struct {
	Code int64  //http code
	Body []byte // response body
}
