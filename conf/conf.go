package conf

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wuqinqiang/easycar/core/notify"

	"github.com/wuqinqiang/easycar/core/notify/dingtalk"

	"github.com/wuqinqiang/easycar/core/notify/telegram"

	"github.com/wuqinqiang/easycar/core/notify/lark"

	"github.com/hashicorp/consul/api"

	"github.com/wuqinqiang/easycar/core/registry/consulx"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"github.com/wuqinqiang/easycar/core/transport/common"

	"github.com/wuqinqiang/easycar/tracing"

	"github.com/wuqinqiang/easycar/core/servers/grpcsrv"

	"github.com/wuqinqiang/easycar/core/dao/mongodb"

	gormx "github.com/wuqinqiang/easycar/core/dao/gorm"

	"github.com/wuqinqiang/easycar/core/registry"

	"github.com/wuqinqiang/easycar/core/registry/etcdx"

	_ "net/http/pprof"
)

type (
	Mode string
)

const (
	File Mode = "file"
	Etcd Mode = "etcd"
	Env  Mode = "env"
	//Add more conf schema here
)

type (
	Settings struct {
		Server              `yaml:"server"`
		DB                  DB               `yaml:"db"`
		Timeout             int64            `yaml:"timeout"`
		AutomaticExecution2 bool             `yaml:"automaticExecution2"`
		Tracing             Tracing          `yaml:"tracing"`
		Registry            RegistrySettings `yaml:"registry"`
		Cron                Cron             `yaml:"cron"`
		Notify              Notify           `yaml:"notify"`
	}

	//DB Config
	DB struct {
		Driver  string           `yaml:"driver"`
		Mysql   gormx.Settings   `yaml:"mysql"`
		Mongodb mongodb.Settings `yaml:"mongodb"`
	}

	//Notify Config
	Notify struct {
		Lark     lark.NotifyConfig     `yaml:"lark"`
		Tg       telegram.NotifyConfig `yaml:"tg"`
		Dingtalk dingtalk.NotifyConfig `yaml:"dingtalk"`
	}

	Server struct {
		Http httpsrv.Http `yaml:"http"`
		Grpc grpcsrv.Grpc `yaml:"grpc"`
	}

	RegistrySettings struct {
		Etcd   etcdx.Conf   `yaml:"etcd"`
		Consul consulx.Conf `yaml:"consul"`
	}

	Tracing struct {
		JaegerUri string `yaml:"jaegerUrl"`
	}
	Cron struct {
		MaxTimes     int `yaml:"maxTimes"`
		TimeInterval int `yaml:"timeInterval"`
	}
)

type Conf interface {
	Load() (*Settings, error)
}

func (db *DB) Init() {
	switch db.Driver {
	case "mysql":
		db.Mysql.Init()
	case "mongodb":
		db.Mongodb.Init()
	default:
		panic(fmt.Errorf("no support %s database", db.Driver))
	}
}

func (s *Settings) Init() {
	s.DB.Init()
	tracing.Init(s.Tracing.JaegerUri)

	// todo custom port
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	if s.Http.ListenOn == "" {
		s.Http.ListenOn = "0.0.0.0:8085"
	}
	if s.Grpc.ListenOn == "" {
		s.Grpc.ListenOn = "0.0.0.0:8088"
	}
	if s.Timeout > 0 {
		common.ReplaceTimeout(time.Duration(s.Timeout) * time.Second)
	}
}

func (s *Settings) SetRegistry() bool {
	return !s.Registry.Etcd.Empty() || !s.Registry.Consul.Empty()
	// todo add more registry center
}
func (s *Settings) GetRegistry() (registry.Registry, error) {
	if !s.Registry.Etcd.Empty() {
		return etcdx.New(s.Registry.Etcd)
	}

	// consul and add others?
	client, err := api.NewClient(s.Registry.Consul.Conf())
	if err != nil {
		return nil, err
	}
	return consulx.New(client), nil
}

func (n *Notify) Senders() (sender []notify.Sender) {
	if n.Tg.Token != "" && n.Tg.ChatID != "" {
		sender = append(sender, n.Tg)
	}
	if n.Lark.WebhookURL != "" {
		sender = append(sender, n.Lark)
	}
	if n.Dingtalk.SignSecret != "" && n.Dingtalk.WebhookURL != "" {
		sender = append(sender, n.Dingtalk)
	}
	return
}
