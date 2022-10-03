package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
)

type BranchImpl struct {
}

func (b BranchImpl) GetCollection() *mongo.Collection {
	return database.Collection("branch")
}

func (b BranchImpl) CreateBatches(ctx context.Context, list entity.BranchList) error {
	var (
		data []interface{}
	)

	for i := range list {
		data = append(data, bson.D{
			{Key: "branch_id", Value: list[i].BranchId},
			{Key: "url", Value: list[i].Url},
			{Key: "req_data", Value: list[i].ReqData},
			{Key: "tran_type", Value: list[i].TranType},
			{Key: "protocol", Value: list[i].Protocol},
			{Key: "action", Value: list[i].Action},
			{Key: "state", Value: list[i].State},
			{Key: "level", Value: list[i].Level},
			{Key: "req_header", Value: list[i].ReqHeader},
			{Key: "timeout", Value: list[i].Timeout},
			{Key: "g_id", Value: list[i].GID},
			{Key: "create_time", Value: list[i].CreateTime},
			{Key: "update_time", Value: list[i].UpdateTime},
		})
	}
	_, err := b.GetCollection().InsertMany(ctx, data)
	return err
}

func (b BranchImpl) GetBranches(ctx context.Context, gid string) (list entity.BranchList, err error) {
	filter := bson.D{{Key: "g_id", Value: gid}}

	var (
		cur *mongo.Cursor
	)
	cur, err = b.GetCollection().Find(ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(ctx, &list)
	return
}

func (b BranchImpl) UpdateBranchStateByGid(ctx context.Context, branchId string, state consts.BranchState, errMsg string) (int64, error) {
	fmt.Println(branchId, state)
	filter := bson.D{{Key: "branch_id", Value: branchId}}
	updates := bson.M{"$set": bson.M{"state": state, "last_err_msg": errMsg, "update_time": time.Now().Unix()}}
	result, err := b.GetCollection().UpdateOne(ctx, filter, updates)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, err
}
