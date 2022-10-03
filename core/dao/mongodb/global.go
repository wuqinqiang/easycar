package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
)

type GlobalImpl struct {
}

func (g GlobalImpl) GetCollection() *mongo.Collection {
	return database.Collection("global")
}

func (g GlobalImpl) CreateGlobal(ctx context.Context, global *entity.Global) (err error) {
	_, err = g.GetCollection().InsertOne(ctx, global)
	return
}

func (g GlobalImpl) GetGlobal(ctx context.Context, gid string) (global entity.Global, err error) {
	filter := bson.D{{Key: "g_id", Value: gid}}
	err = g.GetCollection().FindOne(ctx, filter).Decode(&global)
	return
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state consts.GlobalState) (int64, error) {
	filter := bson.D{{Key: "g_id", Value: gid}}
	var endTime int64 = 0

	if state == consts.Committed || state == consts.Rollbacked {
		endTime = time.Now().Unix()
	}
	updates := bson.M{"$set": bson.M{"state": state, "end_time": endTime, "update_time": time.Now().Unix()}}
	result, err := g.GetCollection().UpdateOne(ctx, filter, updates)
	return result.ModifiedCount, err
}
