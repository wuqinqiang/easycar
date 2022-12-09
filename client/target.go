package client

import "fmt"

// BuildDirectTarget covert uri to target
// uri:127.0.0.1:8089,127.0.0.1:8085 => direct:///127.0.0.1:8089,127.0.0.1:8085
func BuildDirectTarget(uri string) string {
	return fmt.Sprintf("direct:///%s", uri)
}

//BuildDiscoveryTarget
func BuildDiscoveryTarget(uri string) string {
	return fmt.Sprintf("discovery:///%s", uri)
}
