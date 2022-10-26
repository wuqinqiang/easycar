package grpcsrv

type Grpc struct {
	ListenOn string  `yaml:"listenOn"`
	KeyFile  string  `yaml:"keyFile"`
	CertFile string  `yaml:"certFile"`
	Gateway  Gateway `yaml:"gateway"`
}
type Gateway struct {
	IsOpen     bool   `yaml:"isOpen"`
	CertFile   string `yaml:"certFile"`
	ServerName string `yaml:"serverName"`
}

func (grpc *Grpc) Tls() bool {
	return grpc.KeyFile != "" && grpc.CertFile != ""
}

func (grpc *Grpc) IsOpenGateway() bool {
	return grpc.Gateway.IsOpen
}
