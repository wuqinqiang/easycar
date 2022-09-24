package conf

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
	MaxLifetime  int    `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Mysql) Init() *gorm.DB {
	var (
		err error
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
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
	// hooks
	//db.Callback().Query().After("gorm:query").Register("afterQuery", registerCallback) //nolint:errcheck
	//
	//db.Callback().Create().After("gorm:create").Register("afterCreate", registerCallback) //nolint:errcheck
	//
	//db.Callback().Update().After("gorm:update").Register("afterUpdate", registerCallback) //nolint:errcheck
	return db
}

//func registerCallback(db *gorm.DB) {
//	_, span := tracing.Tracer().Start(db.Statement.Context, db.Statement.Table)
//	span.SetAttributes(
//		attribute.String("sql", db.Statement.SQL.String()),
//		attribute.String("row", strconv.Itoa(int(db.RowsAffected))),
//	)
//	span.End()
//}

func GetDb() *gorm.DB {
	if db == nil {
		panic("db is nil")
	}
	return db
}
