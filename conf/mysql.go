package conf

import (
	"time"

	"github.com/wuqinqiang/easycar/core/dao"

	gormDb "github.com/wuqinqiang/easycar/core/dao/gorm"

	"github.com/wuqinqiang/easycar/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlSettings conf
type MysqlSettings struct {
	DbURL        string `json:"dbURL" yaml:"dbURL"`
	MaxLifetime  int    `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *MysqlSettings) Init() {
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

	dao.SetTransaction(gormDb.NewDao(db))

}
