package conf

type Conf struct {
	Server Server
	Mysql  Mysql
}

// Server conf
type Server struct {
	Addr string `json:"addr" yaml:"addr"`
}

// Mysql conf
type Mysql struct {
	Url           string `json:"url" yaml:"url"`
	DbName        string `json:"dbName" yaml:"dbName"`
	Port          int32  `json:"port" yaml:"dbName"`
	User          string `json:"user" yaml:"user"`
	Password      string `json:"password" yaml:"password"`
	MaxLifetime   int64  `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleNumber int64  `json:"maxIdleNumber" yaml:"maxIdleNumber"`
	MaxOpenNumber int64  `json:"maxOpenNumber" yaml:"maxOpenNumber"`
}
