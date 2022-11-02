package balancer

import (
	"github.com/zehuamama/balancer/balancer"
	grpcBalancer "google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

var (
	BuilderName                     = "easycarBalancer"
	_           base.PickerBuilder  = (*Builder)(nil)
	_           grpcBalancer.Picker = (*Picker)(nil)
)

func init() {
	builder := base.NewBalancerBuilder(
		BuilderName,
		&Builder{},
		base.Config{HealthCheck: true})
	grpcBalancer.Register(builder)
}

type Builder struct{}

func (b *Builder) Build(info base.PickerBuildInfo) grpcBalancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(grpcBalancer.ErrNoSubConnAvailable)
	}
	var (
		hosts []string
	)
	picker := &Picker{
		readyAddrConn: make(map[string]grpcBalancer.SubConn),
	}
	for subConn, connInfo := range info.ReadySCs {
		hosts = append(hosts, connInfo.Address.Addr)
		picker.readyAddrConn[connInfo.Address.Addr] = subConn
	}
	balancer, err := balancer.Build(balancer.P2CBalancer, hosts)
	if err != nil {
		panic(err)
	}
	picker.balancer = balancer
	return picker
}

type Picker struct {
	balancer      balancer.Balancer
	readyAddrConn map[string]grpcBalancer.SubConn
}

func (p *Picker) Pick(info grpcBalancer.PickInfo) (grpcBalancer.PickResult, error) {
	// todo get client ip
	// todo Add middleware
	addr, err := p.balancer.Balance(info.FullMethodName)
	if err != nil {
		return grpcBalancer.PickResult{}, grpcBalancer.ErrNoSubConnAvailable
	}
	subConn, ok := p.readyAddrConn[addr]
	if !ok {
		return grpcBalancer.PickResult{}, grpcBalancer.ErrNoSubConnAvailable
	}
	return grpcBalancer.PickResult{
		SubConn: subConn,
		Done: func(info grpcBalancer.DoneInfo) {
			p.balancer.Done(addr)
		},
	}, nil
}
