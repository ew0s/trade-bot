package postgres

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

const QueryBuilderDialect = "postgres"

//go:generate mockgen -source=./query_builder.go -destination=./mock/query_builder.go -package=mock
type QueryBuilder interface {
	From(cols ...interface{}) *goqu.SelectDataset
	Select(cols ...interface{}) *goqu.SelectDataset
	Update(table interface{}) *goqu.UpdateDataset
	Insert(table interface{}) *goqu.InsertDataset
	Delete(table interface{}) *goqu.DeleteDataset
	Truncate(table ...interface{}) *goqu.TruncateDataset
}

func NewBuilder(db *sql.DB) QueryBuilder {
	return goqu.New(QueryBuilderDialect, db)
}

func NewTxBuilder(tx *sql.Tx) QueryBuilder {
	return goqu.NewTx(QueryBuilderDialect, tx)
}
