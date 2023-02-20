// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/HelliWrold1/quaver/dal/model"
)

func newLike(db *gorm.DB, opts ...gen.DOOption) like {
	_like := like{}

	_like.likeDo.UseDB(db, opts...)
	_like.likeDo.UseModel(&model.Like{})

	tableName := _like.likeDo.TableName()
	_like.ALL = field.NewAsterisk(tableName)
	_like.ID = field.NewInt64(tableName, "id")
	_like.VideoID = field.NewInt64(tableName, "video_id")
	_like.LikerID = field.NewInt64(tableName, "liker_id")
	_like.Delete_ = field.NewInt64(tableName, "delete")

	_like.fillFieldMap()

	return _like
}

type like struct {
	likeDo

	ALL     field.Asterisk
	ID      field.Int64
	VideoID field.Int64
	LikerID field.Int64
	Delete_ field.Int64

	fieldMap map[string]field.Expr
}

func (l like) Table(newTableName string) *like {
	l.likeDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l like) As(alias string) *like {
	l.likeDo.DO = *(l.likeDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *like) updateTableName(table string) *like {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.VideoID = field.NewInt64(table, "video_id")
	l.LikerID = field.NewInt64(table, "liker_id")
	l.Delete_ = field.NewInt64(table, "delete")

	l.fillFieldMap()

	return l
}

func (l *like) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *like) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["id"] = l.ID
	l.fieldMap["video_id"] = l.VideoID
	l.fieldMap["liker_id"] = l.LikerID
	l.fieldMap["delete"] = l.Delete_
}

func (l like) clone(db *gorm.DB) like {
	l.likeDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l like) replaceDB(db *gorm.DB) like {
	l.likeDo.ReplaceDB(db)
	return l
}

type likeDo struct{ gen.DO }

type ILikeDo interface {
	gen.SubQuery
	Debug() ILikeDo
	WithContext(ctx context.Context) ILikeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILikeDo
	WriteDB() ILikeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILikeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILikeDo
	Not(conds ...gen.Condition) ILikeDo
	Or(conds ...gen.Condition) ILikeDo
	Select(conds ...field.Expr) ILikeDo
	Where(conds ...gen.Condition) ILikeDo
	Order(conds ...field.Expr) ILikeDo
	Distinct(cols ...field.Expr) ILikeDo
	Omit(cols ...field.Expr) ILikeDo
	Join(table schema.Tabler, on ...field.Expr) ILikeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILikeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILikeDo
	Group(cols ...field.Expr) ILikeDo
	Having(conds ...gen.Condition) ILikeDo
	Limit(limit int) ILikeDo
	Offset(offset int) ILikeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILikeDo
	Unscoped() ILikeDo
	Create(values ...*model.Like) error
	CreateInBatches(values []*model.Like, batchSize int) error
	Save(values ...*model.Like) error
	First() (*model.Like, error)
	Take() (*model.Like, error)
	Last() (*model.Like, error)
	Find() ([]*model.Like, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Like, err error)
	FindInBatches(result *[]*model.Like, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Like) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILikeDo
	Assign(attrs ...field.AssignExpr) ILikeDo
	Joins(fields ...field.RelationField) ILikeDo
	Preload(fields ...field.RelationField) ILikeDo
	FirstOrInit() (*model.Like, error)
	FirstOrCreate() (*model.Like, error)
	FindByPage(offset int, limit int) (result []*model.Like, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILikeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l likeDo) Debug() ILikeDo {
	return l.withDO(l.DO.Debug())
}

func (l likeDo) WithContext(ctx context.Context) ILikeDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l likeDo) ReadDB() ILikeDo {
	return l.Clauses(dbresolver.Read)
}

func (l likeDo) WriteDB() ILikeDo {
	return l.Clauses(dbresolver.Write)
}

func (l likeDo) Session(config *gorm.Session) ILikeDo {
	return l.withDO(l.DO.Session(config))
}

func (l likeDo) Clauses(conds ...clause.Expression) ILikeDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l likeDo) Returning(value interface{}, columns ...string) ILikeDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l likeDo) Not(conds ...gen.Condition) ILikeDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l likeDo) Or(conds ...gen.Condition) ILikeDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l likeDo) Select(conds ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l likeDo) Where(conds ...gen.Condition) ILikeDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l likeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILikeDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l likeDo) Order(conds ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l likeDo) Distinct(cols ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l likeDo) Omit(cols ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l likeDo) Join(table schema.Tabler, on ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l likeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILikeDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l likeDo) RightJoin(table schema.Tabler, on ...field.Expr) ILikeDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l likeDo) Group(cols ...field.Expr) ILikeDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l likeDo) Having(conds ...gen.Condition) ILikeDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l likeDo) Limit(limit int) ILikeDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l likeDo) Offset(offset int) ILikeDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l likeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILikeDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l likeDo) Unscoped() ILikeDo {
	return l.withDO(l.DO.Unscoped())
}

func (l likeDo) Create(values ...*model.Like) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l likeDo) CreateInBatches(values []*model.Like, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l likeDo) Save(values ...*model.Like) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l likeDo) First() (*model.Like, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Like), nil
	}
}

func (l likeDo) Take() (*model.Like, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Like), nil
	}
}

func (l likeDo) Last() (*model.Like, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Like), nil
	}
}

func (l likeDo) Find() ([]*model.Like, error) {
	result, err := l.DO.Find()
	return result.([]*model.Like), err
}

func (l likeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Like, err error) {
	buf := make([]*model.Like, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l likeDo) FindInBatches(result *[]*model.Like, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l likeDo) Attrs(attrs ...field.AssignExpr) ILikeDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l likeDo) Assign(attrs ...field.AssignExpr) ILikeDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l likeDo) Joins(fields ...field.RelationField) ILikeDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l likeDo) Preload(fields ...field.RelationField) ILikeDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l likeDo) FirstOrInit() (*model.Like, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Like), nil
	}
}

func (l likeDo) FirstOrCreate() (*model.Like, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Like), nil
	}
}

func (l likeDo) FindByPage(offset int, limit int) (result []*model.Like, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l likeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l likeDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l likeDo) Delete(models ...*model.Like) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *likeDo) withDO(do gen.Dao) *likeDo {
	l.DO = *do.(*gen.DO)
	return l
}
