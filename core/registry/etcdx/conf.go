package etcdx

// Conf is the config item with the given key on etcd.
type Conf struct {
	Hosts []string
	Key   string
	User  string `json:"user"`
	Pass  string `json:"pass"`
	//CertFile           string `json:",optional"`
	//CertKeyFile        string `json:",optional=CertFile"`
	//CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool `json:""`
}
