package mongodb

import (
	"github.com/wuqinqiang/easycar/core/dao"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	database *mongo.Database
)

type Dao struct {
	dao.BranchDao
	dao.GlobalDao
}

func NewDao(client *mongo.Client, databaseName string) Dao {
	database = client.Database(databaseName)
	return Dao{
		BranchDao: new(BranchImpl),
		GlobalDao: new(GlobalImpl),
	}
}
