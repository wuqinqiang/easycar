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

// RegisterBalancerWithAlgorithm register grpc balancer with algorithm
func RegisterBalancerWithAlgorithm(name balancer.Algorithm) {
	balancer.Register(balancer.WithAlgorithm(name))
}

// RegisterBuilder register Builder with discovery
func RegisterBuilder(discovery registry.Discovery) {
	resolver.Register(discovery)
}
