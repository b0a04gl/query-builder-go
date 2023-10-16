package querybuilder

import (
	"testing"
	"strings"
)

func TestSQLBuilder_Select(t *testing.T) {
	builder := NewSQLBuilder()
	builder.Select("id", "name")

	query, _ := builder.Build()
	expectedQuery := "SELECT id, name"

	if query != expectedQuery {
		t.Error("expected and actual output not matching")
	}
}

func TestSQLBuilder_From(t *testing.T) {
	builder := NewSQLBuilder()
	builder.From("players")

	query, _ := builder.Build()
	expectedQuery := " FROM players"

	if !strings.Contains(query, expectedQuery) {
		t.Error("generated query doesn't contain expected clauses")
	}
}

func TestSQLBuilder_Where(t *testing.T) {
	builder := NewSQLBuilder()
	builder.Where("total_runs > ?", 2500)

	query, args := builder.Build()
	expectedQuery := " WHERE total_runs > ?"
	expectedArgs := []interface{}{2500}

	if !strings.Contains(query, expectedQuery) {
		t.Error("generated query doesn't contain expected clauses")
	}
	if len(args) != len(expectedArgs) || args[0] != expectedArgs[0] {
		t.Error("argument parameters don't match")
	}
}

func TestSQLBuilder_OrderBy(t *testing.T) {
	builder := NewSQLBuilder()
	builder.OrderBy("total_wickets DESC")

	query, _ := builder.Build()
	expectedQuery := " ORDER BY total_wickets DESC"

	if !strings.Contains(query, expectedQuery) {
		t.Error("generated query doesn't contain expected clauses")
	}
}

func TestSQLBuilder_Build(t *testing.T) {
	builder := NewSQLBuilder()
	query, args := builder.
		Select("id", "name").
		From("players").
		Where("total_runs > ?", 2500).
		OrderBy("total_wickets DESC").
		Build()

	expectedQuery := "SELECT id, name FROM players WHERE total_runs > ? ORDER BY total_wickets DESC"
	expectedArgs := []interface{}{2500}

	if query != expectedQuery {
		t.Error("generated query doesn't contain expected clauses")
	}
	if len(args) != len(expectedArgs) || args[0] != expectedArgs[0] {
		t.Error("argument parameters don't match")
	}
}
