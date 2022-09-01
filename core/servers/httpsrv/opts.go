package httpsrv

import "time"

type Opt func(srv *HttpSrv)

func WithTimeout(t time.Duration) Opt {
	return func(srv *HttpSrv) {
		srv.timeout = t
	}
}
