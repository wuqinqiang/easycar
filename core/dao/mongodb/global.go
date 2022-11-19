package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
)

type GlobalImpl struct {
}

func (g GlobalImpl) IncrTryTimes(ctx context.Context, gid string, nextCronTime int) error {
	filter := bson.D{{Key: "g_id", Value: gid}}
	updates := bson.M{"$set": bson.M{"next_cron_time": nextCronTime}, "$inc": bson.M{
		"try_times": 1,
	}}
	_, err := g.GetCollection().UpdateOne(ctx, filter, updates)
	return err
}

func (g GlobalImpl) GetCollection() *mongo.Collection {
	return database.Collection("global")
}

func (g GlobalImpl) CreateGlobal(ctx context.Context, global *entity.Global) (err error) {
	_, err = g.GetCollection().InsertOne(ctx, global)
	return
}

func (g GlobalImpl) FindProcessingList(ctx context.Context, limit, maxTimes int) (list []*entity.Global, err error) {
	now := time.Now()
	before := now.Add(time.Hour * -2)

	var (
		state []string
	)
	state = append(state, consts.P1InProgressStates...)
	state = append(state, consts.P2InProgressStates...)

	filter := bson.M{
		"$and": []bson.M{
			{
				"next_cron_time": bson.M{"$gte": before.Unix(), "$lte": now.Unix()},
			},
			{
				"state": bson.M{"$in": state},
			},
			{
				"try_times": bson.M{"$lt": maxTimes},
			},
		},
	}
	var (
		cur *mongo.Cursor
	)

	cur, err = g.GetCollection().Find(ctx, filter,
		options.Find().SetSort(bson.M{"update_time": 1}).SetLimit(int64(limit)))
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
