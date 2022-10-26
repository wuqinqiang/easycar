package gorm

import (
	"time"

	"github.com/wuqinqiang/easycar/core/dao"
	"github.com/wuqinqiang/easycar/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Settings conf
type Settings struct {
	DbURL        string `json:"dbURL" yaml:"dbURL"`
	MaxLifetime  int    `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Settings) Init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               m.DbURL,
		DefaultStringSize: 256,
	}), &gorm.Config{SkipDefaultTransaction: true})
	tools.ErrToPanic(err)

	d, err := db.DB()

	tools.ErrToPanic(err)

	if m.MaxLifetime > 0 {
		d.SetConnMaxLifetime(time.Duration(m.MaxLifetime) * time.Second)
	}
	if m.MaxOpenConns > 0 {
		d.SetMaxOpenConns(m.MaxOpenConns)
	}

	if m.MaxIdleConns > 0 {
		d.SetMaxIdleConns(m.MaxIdleConns)
	}

	dao.SetTransaction(NewDao(db))
}
