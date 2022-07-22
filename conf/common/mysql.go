package common

import (
	"time"

	"github.com/wuqinqiang/easycar/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Mysql conf
type Mysql struct {
	DbURL        string `json:"dbURL" yaml:"dbURL"`
	MaxLifetime  int64  `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleConns int64  `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int64  `json:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Mysql) Init() *gorm.DB {
	var (
		err error
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               m.DbURL,
		DefaultStringSize: 256,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}

	d, err := db.DB()
	if err != nil {
		panic(err)
	}

	if m.MaxLifetime > 0 {
		d.SetConnMaxLifetime(time.Duration(m.MaxLifetime) * time.Second)
	}
	if m.MaxOpenConns > 0 {
		d.SetMaxOpenConns(int(m.MaxOpenConns))
	}

	if m.MaxIdleConns > 0 {
		d.SetMaxIdleConns(int(m.MaxIdleConns))
	}

	if err != nil {
		tools.ErrToPanic(err)
	}
	return db
}

func GetDb() *gorm.DB {
	if db == nil {
		panic("db is nil")
	}
	return db
}
