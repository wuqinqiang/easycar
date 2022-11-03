package balancer

// TacticsName is tactics of balancer
type TacticsName int

const (
	IPHashBalancer TacticsName = iota + 1
	ConsistentHashBalancer
	P2CBalancer
	RandomBalancer
	R2Balancer
	LeastLoadBalancer
	BoundedBalancer
)

func (l TacticsName) Name() string {
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
