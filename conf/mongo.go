package conf

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout = time.Second * 5
	MinPool        = 10
	DefaultMaxPool = 100
	AllowMaxPool   = 250
)

type MongoSetting struct {
	Uri     string `yaml:"uri"` // example: mongodb://localhost:27017
	MinPool int    `yaml:"minPool"`
	MaxPool int    `yaml:"maxPool"`
}

func (settings *MongoSetting) GetClient() (*mongo.Client, error) {
	if settings.Uri == "" {
		return nil, fmt.Errorf("empty mongo uri")
	}
	opts := new(options.ClientOptions)
	opts.SetConnectTimeout(defaultTimeout)
	opts.SetServerSelectionTimeout(defaultTimeout)
	opts.SetMinPoolSize(uint64(settings.MinPool))
	maxPool := settings.MaxPool
	if maxPool > AllowMaxPool {
		maxPool = AllowMaxPool
	}
	if maxPool < settings.MinPool {
		maxPool = DefaultMaxPool
	}
	opts.SetMaxPoolSize(uint64(maxPool))

	opts.ApplyURI(settings.Uri)
	var (
		err    error
		client *mongo.Client
	)
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect err: %v", err)
	}
	// ping client conn
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo err:%v", err)
	}
	return client, nil
}
