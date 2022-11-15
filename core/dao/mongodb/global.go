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

func (g GlobalImpl) FindProcessingList(ctx context.Context, limit int) (list []*entity.Global, err error) {
	now := time.Now()
	before := now.Add(time.Minute * -2)

	filter := bson.M{
		"create_time": bson.M{"$gte": before.Unix(), "$lte": now.Unix()},
		"state": bson.M{"$in": []string{string(consts.Phase1Preparing),
			string(consts.Phase2Committing), string(consts.Phase2Rollbacking)}},
	}
	var (
		cur *mongo.Cursor
	)
	cur, err = g.GetCollection().Find(ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(ctx, &list)
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
