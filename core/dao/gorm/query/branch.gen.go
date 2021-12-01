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

func newBranch(db *gorm.DB) branch {
	_branch := branch{}

	_branch.branchDo.UseDB(db)
	_branch.branchDo.UseModel(&model.Branch{})

	tableName := _branch.branchDo.TableName()
	_branch.ALL = field.NewField(tableName, "*")
	_branch.ID = field.NewInt32(tableName, "id")
	_branch.Gid = field.NewString(tableName, "gid")
	_branch.URL = field.NewString(tableName, "url")
	_branch.ReqData = field.NewString(tableName, "req_data")
	_branch.BranchID = field.NewString(tableName, "branch_id")
	_branch.BranchType = field.NewString(tableName, "branch_type")
	_branch.State = field.NewString(tableName, "state")
	_branch.FinishTime = field.NewTime(tableName, "finish_time")
	_branch.CreateTime = field.NewTime(tableName, "create_time")
	_branch.UpdateTime = field.NewTime(tableName, "update_time")

	_branch.fieldMap = make(map[string]field.Expr, 10)
	_branch.fieldMap["id"] = _branch.ID
	_branch.fieldMap["gid"] = _branch.Gid
	_branch.fieldMap["url"] = _branch.URL
	_branch.fieldMap["req_data"] = _branch.ReqData
	_branch.fieldMap["branch_id"] = _branch.BranchID
	_branch.fieldMap["branch_type"] = _branch.BranchType
	_branch.fieldMap["state"] = _branch.State
	_branch.fieldMap["finish_time"] = _branch.FinishTime
	_branch.fieldMap["create_time"] = _branch.CreateTime
	_branch.fieldMap["update_time"] = _branch.UpdateTime

	return _branch
}

type branch struct {
	branchDo branchDo

	ALL        field.Field
	ID         field.Int32
	Gid        field.String
	URL        field.String
	ReqData    field.String
	BranchID   field.String
	BranchType field.String
	State      field.String
	FinishTime field.Time
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (b *branch) WithContext(ctx context.Context) *branchDo { return b.branchDo.WithContext(ctx) }

func (b branch) TableName() string { return b.branchDo.TableName() }

func (b *branch) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := b.fieldMap[fieldName]
	return field, ok
}

func (b branch) clone(db *gorm.DB) branch {
	b.branchDo.ReplaceDB(db)
	return b
}

type branchDo struct{ gen.DO }

func (b branchDo) Debug() *branchDo {
	return b.withDO(b.DO.Debug())
}

func (b branchDo) WithContext(ctx context.Context) *branchDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b branchDo) Clauses(conds ...clause.Expression) *branchDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b branchDo) Not(conds ...gen.Condition) *branchDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b branchDo) Or(conds ...gen.Condition) *branchDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b branchDo) Select(conds ...field.Expr) *branchDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b branchDo) Where(conds ...gen.Condition) *branchDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b branchDo) Order(conds ...field.Expr) *branchDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b branchDo) Distinct(cols ...field.Expr) *branchDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b branchDo) Omit(cols ...field.Expr) *branchDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b branchDo) Join(table schema.Tabler, on ...field.Expr) *branchDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b branchDo) LeftJoin(table schema.Tabler, on ...field.Expr) *branchDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b branchDo) RightJoin(table schema.Tabler, on ...field.Expr) *branchDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b branchDo) Group(cols ...field.Expr) *branchDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b branchDo) Having(conds ...gen.Condition) *branchDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b branchDo) Limit(limit int) *branchDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b branchDo) Offset(offset int) *branchDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b branchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *branchDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b branchDo) Unscoped() *branchDo {
	return b.withDO(b.DO.Unscoped())
}

func (b branchDo) Create(values ...*model.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b branchDo) CreateInBatches(values []*model.Branch, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b branchDo) Save(values ...*model.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b branchDo) First() (*model.Branch, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Take() (*model.Branch, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Last() (*model.Branch, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Find() ([]*model.Branch, error) {
	result, err := b.DO.Find()
	return result.([]*model.Branch), err
}

func (b branchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Branch, error) {
	result, err := b.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Branch), err
}

func (b branchDo) FindInBatches(result []*model.Branch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(&result, batchSize, fc)
}

func (b branchDo) Attrs(attrs ...field.AssignExpr) *branchDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b branchDo) Assign(attrs ...field.AssignExpr) *branchDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b branchDo) Joins(field field.RelationField) *branchDo {
	return b.withDO(b.DO.Joins(field))
}

func (b branchDo) Preload(field field.RelationField) *branchDo {
	return b.withDO(b.DO.Preload(field))
}

func (b branchDo) FirstOrInit() (*model.Branch, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) FirstOrCreate() (*model.Branch, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) FindByPage(offset int, limit int) (result []*model.Branch, count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	result, err = b.Offset(offset).Limit(limit).Find()
	return
}

func (b *branchDo) withDO(do gen.Dao) *branchDo {
	b.DO = *do.(*gen.DO)
	return b
}
