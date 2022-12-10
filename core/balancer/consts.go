package balancer

// Algorithm is algorithm of balancer
type Algorithm int

const (
	IPHashBalancer Algorithm = iota + 1
	ConsistentHashBalancer
	P2CBalancer
	RandomBalancer
	R2Balancer
	LeastLoadBalancer
	BoundedBalancer
)

func (l Algorithm) Name() string {
	switch l {
	case IPHashBalancer:
		return "ip-hash"
	case ConsistentHashBalancer:
		return "consistent-hash"
	case P2CBalancer:
		return "p2c"
	case RandomBalancer:
		return "random"
	case R2Balancer:
		return "round-robin"
	case LeastLoadBalancer:
		return "least-load"
	case BoundedBalancer:
		return "bounded"
	default:
		return ""
	}
}
