package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/wuqinqiang/easycar/core/dao"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout = time.Second * 5
	MinPool        = 10
	DefaultMaxPool = 100
	AllowMaxPool   = 250
)

type Settings struct {
	Uri     string `yaml:"url"` // example: mongodb://localhost:27017
	MinPool int    `yaml:"minPool"`
	MaxPool int    `yaml:"maxPool"`
}

func (settings *Settings) Init() {
	if settings.Uri == "" {
		panic(fmt.Errorf("emtpy mongo uri"))
	}
	opts := new(options.ClientOptions)
	opts.SetConnectTimeout(defaultTimeout)
	opts.SetServerSelectionTimeout(defaultTimeout)
	minPool := settings.MinPool
	if minPool < MinPool {
		minPool = MinPool
	}
	opts.SetMinPoolSize(uint64(minPool))
	maxPool := settings.MaxPool
	if maxPool < settings.MinPool {
		maxPool = DefaultMaxPool
	}

	if maxPool > AllowMaxPool {
		maxPool = AllowMaxPool
	}

	opts.SetMaxPoolSize(uint64(maxPool))

	opts.ApplyURI(settings.Uri)
	var (
		err    error
		client *mongo.Client
	)
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(fmt.Errorf("mongo connect err: %v", err))
	}
	// ping client conn
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Errorf("ping mongo err:%v", err))
	}
	// todo option
	dao.SetTransaction(NewDao(client, "easycar"))
}
