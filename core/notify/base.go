package notify

import "fmt"

type Content struct {
	gid string
	err error
}

func NewContext(gid string, err error) Content {
	return Content{
		gid: gid,
		err: err,
	}
}

func (context *Content) Msg() string {
	return fmt.Sprintf("gid:%s err:%v", context.gid, context.err)
}
