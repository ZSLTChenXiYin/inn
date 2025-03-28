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

	"file/biz/dal/model"
)

func newFile(db *gorm.DB, opts ...gen.DOOption) file {
	_file := file{}

	_file.fileDo.UseDB(db, opts...)
	_file.fileDo.UseModel(&model.File{})

	tableName := _file.fileDo.TableName()
	_file.ALL = field.NewAsterisk(tableName)
	_file.FileID = field.NewString(tableName, "file_id")
	_file.FileName = field.NewString(tableName, "file_name")
	_file.FilePath = field.NewString(tableName, "file_path")
	_file.FileSize = field.NewInt64(tableName, "file_size")
	_file.OwnerID = field.NewString(tableName, "owner_id")
	_file.AccessLevel = field.NewInt32(tableName, "access_level")
	_file.FilePasswordHash = field.NewString(tableName, "file_password_hash")
	_file.CheckSum = field.NewString(tableName, "check_sum")
	_file.CreatedAt = field.NewTime(tableName, "created_at")
	_file.UpdatedAt = field.NewTime(tableName, "updated_at")
	_file.DeletedAt = field.NewField(tableName, "deleted_at")

	_file.fillFieldMap()

	return _file
}

type file struct {
	fileDo fileDo

	ALL              field.Asterisk
	FileID           field.String
	FileName         field.String
	FilePath         field.String
	FileSize         field.Int64
	OwnerID          field.String
	AccessLevel      field.Int32
	FilePasswordHash field.String
	CheckSum         field.String
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field

	fieldMap map[string]field.Expr
}

func (f file) Table(newTableName string) *file {
	f.fileDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f file) As(alias string) *file {
	f.fileDo.DO = *(f.fileDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *file) updateTableName(table string) *file {
	f.ALL = field.NewAsterisk(table)
	f.FileID = field.NewString(table, "file_id")
	f.FileName = field.NewString(table, "file_name")
	f.FilePath = field.NewString(table, "file_path")
	f.FileSize = field.NewInt64(table, "file_size")
	f.OwnerID = field.NewString(table, "owner_id")
	f.AccessLevel = field.NewInt32(table, "access_level")
	f.FilePasswordHash = field.NewString(table, "file_password_hash")
	f.CheckSum = field.NewString(table, "check_sum")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *file) WithContext(ctx context.Context) *fileDo { return f.fileDo.WithContext(ctx) }

func (f file) TableName() string { return f.fileDo.TableName() }

func (f file) Alias() string { return f.fileDo.Alias() }

func (f file) Columns(cols ...field.Expr) gen.Columns { return f.fileDo.Columns(cols...) }

func (f *file) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *file) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 11)
	f.fieldMap["file_id"] = f.FileID
	f.fieldMap["file_name"] = f.FileName
	f.fieldMap["file_path"] = f.FilePath
	f.fieldMap["file_size"] = f.FileSize
	f.fieldMap["owner_id"] = f.OwnerID
	f.fieldMap["access_level"] = f.AccessLevel
	f.fieldMap["file_password_hash"] = f.FilePasswordHash
	f.fieldMap["check_sum"] = f.CheckSum
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f file) clone(db *gorm.DB) file {
	f.fileDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f file) replaceDB(db *gorm.DB) file {
	f.fileDo.ReplaceDB(db)
	return f
}

type fileDo struct{ gen.DO }

func (f fileDo) Debug() *fileDo {
	return f.withDO(f.DO.Debug())
}

func (f fileDo) WithContext(ctx context.Context) *fileDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileDo) ReadDB() *fileDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileDo) WriteDB() *fileDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileDo) Session(config *gorm.Session) *fileDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileDo) Clauses(conds ...clause.Expression) *fileDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileDo) Returning(value interface{}, columns ...string) *fileDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileDo) Not(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileDo) Or(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileDo) Select(conds ...field.Expr) *fileDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileDo) Where(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileDo) Order(conds ...field.Expr) *fileDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileDo) Distinct(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileDo) Omit(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileDo) Join(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileDo) RightJoin(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileDo) Group(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileDo) Having(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileDo) Limit(limit int) *fileDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileDo) Offset(offset int) *fileDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fileDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileDo) Unscoped() *fileDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileDo) Create(values ...*model.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileDo) CreateInBatches(values []*model.File, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileDo) Save(values ...*model.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileDo) First() (*model.File, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Take() (*model.File, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Last() (*model.File, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Find() ([]*model.File, error) {
	result, err := f.DO.Find()
	return result.([]*model.File), err
}

func (f fileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.File, err error) {
	buf := make([]*model.File, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileDo) FindInBatches(result *[]*model.File, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileDo) Attrs(attrs ...field.AssignExpr) *fileDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileDo) Assign(attrs ...field.AssignExpr) *fileDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileDo) Joins(fields ...field.RelationField) *fileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileDo) Preload(fields ...field.RelationField) *fileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileDo) FirstOrInit() (*model.File, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) FirstOrCreate() (*model.File, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) FindByPage(offset int, limit int) (result []*model.File, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileDo) Delete(models ...*model.File) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileDo) withDO(do gen.Dao) *fileDo {
	f.DO = *do.(*gen.DO)
	return f
}
