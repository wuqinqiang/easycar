package httpsrv

import (
	"context"
	"net/http"
)

type Handler func(ctx context.Context) (http.Handler, error)
