// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newBranch(db *gorm.DB) branch {
	_branch := branch{}

	_branch.branchDo.UseDB(db)
	_branch.branchDo.UseModel(&entity.Branch{})

	tableName := _branch.branchDo.TableName()
	_branch.ALL = field.NewField(tableName, "*")
	_branch.GID = field.NewString(tableName, "g_id")
	_branch.BranchId = field.NewString(tableName, "branch_id")
	_branch.Url = field.NewString(tableName, "url")
	_branch.ReqData = field.NewString(tableName, "req_data")
	_branch.ReqHeader = field.NewString(tableName, "req_header")
	_branch.TranType = field.NewString(tableName, "tran_type")
	_branch.PId = field.NewString(tableName, "p_id")
	_branch.Protocol = field.NewString(tableName, "protocol")
	_branch.Action = field.NewString(tableName, "action")
	_branch.State = field.NewString(tableName, "state")
	_branch.EndTime = field.NewInt64(tableName, "end_time")
	_branch.Level = field.NewUint8(tableName, "level")
	_branch.LastErrMsg = field.NewString(tableName, "last_err_msg")

	_branch.fieldMap = make(map[string]field.Expr, 13)
	_branch.fieldMap["g_id"] = _branch.GID
	_branch.fieldMap["branch_id"] = _branch.BranchId
	_branch.fieldMap["url"] = _branch.Url
	_branch.fieldMap["req_data"] = _branch.ReqData
	_branch.fieldMap["req_header"] = _branch.ReqHeader
	_branch.fieldMap["tran_type"] = _branch.TranType
	_branch.fieldMap["p_id"] = _branch.PId
	_branch.fieldMap["protocol"] = _branch.Protocol
	_branch.fieldMap["action"] = _branch.Action
	_branch.fieldMap["state"] = _branch.State
	_branch.fieldMap["end_time"] = _branch.EndTime
	_branch.fieldMap["level"] = _branch.Level
	_branch.fieldMap["last_err_msg"] = _branch.LastErrMsg

	return _branch
}

type branch struct {
	branchDo branchDo

	ALL        field.Field
	GID        field.String
	BranchId   field.String
	Url        field.String
	ReqData    field.String
	ReqHeader  field.String
	TranType   field.String
	PId        field.String
	Protocol   field.String
	Action     field.String
	State      field.String
	EndTime    field.Int64
	Level      field.Uint8
	LastErrMsg field.String

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

func (b branchDo) Create(values ...*entity.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b branchDo) CreateInBatches(values []*entity.Branch, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b branchDo) Save(values ...*entity.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b branchDo) First() (*entity.Branch, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Branch), nil
	}
}

func (b branchDo) Take() (*entity.Branch, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Branch), nil
	}
}

func (b branchDo) Last() (*entity.Branch, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Branch), nil
	}
}

func (b branchDo) Find() ([]*entity.Branch, error) {
	result, err := b.DO.Find()
	return result.([]*entity.Branch), err
}

func (b branchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*entity.Branch, error) {
	result, err := b.DO.FindInBatch(batchSize, fc)
	return result.([]*entity.Branch), err
}

func (b branchDo) FindInBatches(result []*entity.Branch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
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

func (b branchDo) FirstOrInit() (*entity.Branch, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Branch), nil
	}
}

func (b branchDo) FirstOrCreate() (*entity.Branch, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Branch), nil
	}
}

func (b branchDo) FindByPage(offset int, limit int) (result []*entity.Branch, count int64, err error) {
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
