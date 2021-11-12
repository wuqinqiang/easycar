package mysql

import (
	"encoding/json"

	"fmt"
	"github.com/wuqinqiang/easycar/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

type Config struct {
	Db MysqlConf `json:"db"`
}

type MysqlConf struct {
	Url               string `json:"url"`
	User              string `json:"user"`
	Password          string `json:"password"`
	Port              int64  `json:"port"`
	DbName            string `json:"dbName"`
	MaxLifetime       int64  `json:"maxLifetime"`
	MaxIdleConnNumber int64  `json:"maxIdleNumber"`
	MaxOpenNumber     int64  `json:"maxOpenNumber"`
}

var (
	config      = Config{}
	dbContainer sync.Map
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		utils.ErrToPanic(err)
	}
	bytes, err := os.ReadFile(wd + "/conf/conf.json")
	if err != nil {
		utils.ErrToPanic(err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		utils.ErrToPanic(err)
	}
	fmt.Printf("数据:%+v", config)
}

func NewDb() *gorm.DB {
	dsn := GetDSN()
	value, ok := dbContainer.Load(dsn)
	if ok {
		return value.(*gorm.DB)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               GetDSN(),
		DefaultStringSize: 256,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		utils.ErrToPanic(err)
	}
	dbContainer.Store(dsn, db)
	return db
}

func GetDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Db.User, config.Db.Password, config.Db.Url, config.Db.Port, config.Db.DbName)
}
