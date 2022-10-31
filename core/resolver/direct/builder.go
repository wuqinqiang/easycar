package direct

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(new(Build))
}

type Build struct {
}

func (b Build) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var (
		state resolver.State
	)
	// direct://authority/ip:port,ip:port
	for _, addr := range strings.Split(strings.TrimPrefix(target.URL.Path, "/"), ",") {
		state.Addresses = append(state.Addresses, resolver.Address{Addr: addr})
	}
	if err := cc.UpdateState(state); err != nil {
		return nil, err
	}
	return new(Resolver), nil
}

func (b Build) Scheme() string {
	return "direct"
}
