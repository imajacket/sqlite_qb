package sqlite_qb

import (
	"fmt"
	"strings"
)

type InsertBuilder struct {
	insertStmt string
	columns    string
	values     string
}

func Insert(table string) *InsertBuilder {
	stmt := fmt.Sprintf("INSERT INTO %s", table)
	return &InsertBuilder{
		insertStmt: stmt,
	}
}

func (i *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	i.columns = fmt.Sprintf("(%s)", strings.Join(columns, ","))
	return i
}

func (i *InsertBuilder) Values(values ...any) *InsertBuilder {
	var list []string
	for _, z := range values {
		switch z.(type) {
		case string:
			list = append(list, fmt.Sprintf("'ʭ%sʭ'", z))
		default:
			list = append(list, fmt.Sprintf("ʭ%vʭ", z))
		}
	}

	i.values = fmt.Sprintf("(%s)", strings.Join(list, ","))
	return i
}

func (i *InsertBuilder) ToSql() (string, []interface{}, error) {
	builder := make(map[string]string)

	builder["insert"] = i.insertStmt
	builder["columns"] = i.columns
	builder["values"] = i.values

	return toSQL(builder)
}
