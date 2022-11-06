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

type Option func(options *Options)

type Options struct {
	tactics TacticsName
}

func WithTactics(tacticsName TacticsName) Option {
	return func(options *Options) {
		options.tactics = tacticsName
	}
}

func Register(fns ...Option) {
	// RandomBalancer default tactics
	initOption := &Options{tactics: RandomBalancer}
	for _, fn := range fns {
		fn(initOption)
	}
	builder := &Builder{options: initOption}
	grpcBalancer.Register(base.NewBalancerBuilder(BuilderName, builder, base.Config{HealthCheck: true}))
}

type Builder struct {
	options *Options
}

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
	var (
		err error
	)
	picker.balancer, err = balancer.Build(b.options.tactics.Name(), hosts)
	if err != nil {
		return base.NewErrPicker(err)
	}
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

	p.balancer.Inc(addr)

	return grpcBalancer.PickResult{
		SubConn: subConn,
		Done: func(info grpcBalancer.DoneInfo) {
			p.balancer.Done(addr)
		},
	}, nil
}
