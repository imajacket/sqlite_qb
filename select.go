package sqlite_qb

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	selectStmt string
	fromStmt   string
	whereStmt  string
}

func Select(columns ...string) *SelectBuilder {
	stmt := fmt.Sprintf("SELECT %s", strings.Join(columns, ","))
	return &SelectBuilder{
		selectStmt: stmt,
	}
}

func (s *SelectBuilder) From(table string) *SelectBuilder {
	s.fromStmt = fmt.Sprintf("FROM %s", table)
	return s
}

func (s *SelectBuilder) Where(filter ...string) *SelectBuilder {
	s.whereStmt = fmt.Sprintf("WHERE %s", strings.Join(filter, " AND $"))
	return s
}

func (s *SelectBuilder) ToSQL() (string, []interface{}, error) {
	builder := make(map[string]string)

	if s.selectStmt != "" {
		builder["select"] = s.selectStmt
	}

	if s.fromStmt != "" {
		builder["from"] = s.fromStmt
	}

	if s.whereStmt != "" {
		builder["where"] = s.whereStmt
	}

	return toSQL(builder)
}
