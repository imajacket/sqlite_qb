package sqlite_qb

import (
	"fmt"
	"testing"
)

func TestSelectSimple(t *testing.T) {
	stmt, args, err := Select("*").From("example").ToSQL()
	if err != nil {
		t.Fatalf(`Select("*").From("example").ToSQL() failed - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := fmt.Sprintf("SELECT * FROM example []")

	if result != want {
		t.Fatalf(`Select("*").From("example").ToSQL() failed, want %s, got %s`, want, result)
	}
}

func TestSelectSimpleWhere(t *testing.T) {
	stmt, args, err := Select("*").From("example").Where(Compare(Eq, map[string]interface{}{
		"id": "test_id",
	})).ToSQL()
	if err != nil {
		t.Fatalf(`Select("*").From("example").Where(Compare(Eq, map[string]interface{}{"id": "test_id",})).ToSQL() failed - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := fmt.Sprintf("SELECT * FROM example WHERE 'id' = '?' [test_id]")

	if result != want {
		t.Fatalf(`Select("*").From("example").ToSQL() failed, want %s, got %s`, want, result)
	}
}
