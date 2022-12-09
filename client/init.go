package client

import (
	"github.com/wuqinqiang/easycar/core/balancer"

	"github.com/wuqinqiang/easycar/core/registry"
	"github.com/wuqinqiang/easycar/core/resolver"
)

func init() {
	//default RandomBalancer
	balancer.Register()
}

// RegisterBalancerWithTactics
func RegisterBalancerWithTactics(name balancer.TacticsName) {
	balancer.Register(balancer.WithTactics(name))
}

// RegisterBuilder
func RegisterBuilder(discovery registry.Discovery) {
	resolver.Register(discovery)
}
