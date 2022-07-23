// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newGlobal(db *gorm.DB) global {
	_global := global{}

	_global.globalDo.UseDB(db)
	_global.globalDo.UseModel(&model.Global{})

	tableName := _global.globalDo.TableName()
	_global.ALL = field.NewField(tableName, "*")
	_global.ID = field.NewInt32(tableName, "id")
	_global.GID = field.NewString(tableName, "g_id")
	_global.State = field.NewString(tableName, "state")
	_global.EndTime = field.NewString(tableName, "end_time")
	_global.NextCronTime = field.NewInt32(tableName, "next_cron_time")

	_global.fieldMap = make(map[string]field.Expr, 5)
	_global.fieldMap["id"] = _global.ID
	_global.fieldMap["g_id"] = _global.GID
	_global.fieldMap["state"] = _global.State
	_global.fieldMap["end_time"] = _global.EndTime
	_global.fieldMap["next_cron_time"] = _global.NextCronTime

	return _global
}

type global struct {
	globalDo globalDo

	ALL          field.Field
	ID           field.Int32
	GID          field.String
	State        field.String
	EndTime      field.String
	NextCronTime field.Int32

	fieldMap map[string]field.Expr
}

func (g *global) WithContext(ctx context.Context) *globalDo { return g.globalDo.WithContext(ctx) }

func (g global) TableName() string { return g.globalDo.TableName() }

func (g *global) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g global) clone(db *gorm.DB) global {
	g.globalDo.ReplaceDB(db)
	return g
}

type globalDo struct{ gen.DO }

func (g globalDo) Debug() *globalDo {
	return g.withDO(g.DO.Debug())
}

func (g globalDo) WithContext(ctx context.Context) *globalDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g globalDo) Clauses(conds ...clause.Expression) *globalDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g globalDo) Not(conds ...gen.Condition) *globalDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g globalDo) Or(conds ...gen.Condition) *globalDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g globalDo) Select(conds ...field.Expr) *globalDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g globalDo) Where(conds ...gen.Condition) *globalDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g globalDo) Order(conds ...field.Expr) *globalDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g globalDo) Distinct(cols ...field.Expr) *globalDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g globalDo) Omit(cols ...field.Expr) *globalDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g globalDo) Join(table schema.Tabler, on ...field.Expr) *globalDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g globalDo) LeftJoin(table schema.Tabler, on ...field.Expr) *globalDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g globalDo) RightJoin(table schema.Tabler, on ...field.Expr) *globalDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g globalDo) Group(cols ...field.Expr) *globalDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g globalDo) Having(conds ...gen.Condition) *globalDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g globalDo) Limit(limit int) *globalDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g globalDo) Offset(offset int) *globalDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g globalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *globalDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g globalDo) Unscoped() *globalDo {
	return g.withDO(g.DO.Unscoped())
}

func (g globalDo) Create(values ...*model.Global) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g globalDo) CreateInBatches(values []*model.Global, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g globalDo) Save(values ...*model.Global) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g globalDo) First() (*model.Global, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Global), nil
	}
}

func (g globalDo) Take() (*model.Global, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Global), nil
	}
}

func (g globalDo) Last() (*model.Global, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Global), nil
	}
}

func (g globalDo) Find() ([]*model.Global, error) {
	result, err := g.DO.Find()
	return result.([]*model.Global), err
}

func (g globalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Global, error) {
	result, err := g.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Global), err
}

func (g globalDo) FindInBatches(result []*model.Global, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(&result, batchSize, fc)
}

func (g globalDo) Attrs(attrs ...field.AssignExpr) *globalDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g globalDo) Assign(attrs ...field.AssignExpr) *globalDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g globalDo) Joins(field field.RelationField) *globalDo {
	return g.withDO(g.DO.Joins(field))
}

func (g globalDo) Preload(field field.RelationField) *globalDo {
	return g.withDO(g.DO.Preload(field))
}

func (g globalDo) FirstOrInit() (*model.Global, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Global), nil
	}
}

func (g globalDo) FirstOrCreate() (*model.Global, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Global), nil
	}
}

func (g globalDo) FindByPage(offset int, limit int) (result []*model.Global, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g *globalDo) withDO(do gen.Dao) *globalDo {
	g.DO = *do.(*gen.DO)
	return g
}
