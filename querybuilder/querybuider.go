package querybuilder

import (
	"fmt"
	"strings"
)

type SQLBuilder struct {
	query string
	args  []interface{}
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	b.query = fmt.Sprintf("SELECT %s", strings.Join(columns, ", "))
	return b
}

func (b *SQLBuilder) From(table string) *SQLBuilder {
	b.query += fmt.Sprintf(" FROM %s", table)
	return b
}

func (b *SQLBuilder) Where(condition string, args ...interface{}) *SQLBuilder {
	b.query += " WHERE " + condition
	b.args = append(b.args, args...)
	return b
}

func (b *SQLBuilder) OrderBy(orderBy string) *SQLBuilder {
	b.query += fmt.Sprintf(" ORDER BY %s", orderBy)
	return b
}

func (b *SQLBuilder) Build() (string, []interface{}) {
	return b.query, b.args
}
