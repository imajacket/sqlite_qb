package sqlite_qb

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	stmt, args, err := Insert("example").Columns("Name", "Id").Values(1, "3").ToSql()
	if err != nil {
		t.Fatalf(`Insert("example").Columns("Name", "Id").Values(1, "3").ToSql() - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := fmt.Sprintf("INSERT INTO example (Name,Id) (?,'?') [1 3]")

	if result != want {
		t.Fatalf(`Insert("example").Columns("Name", "Id").Values(1, "3").ToSql() failed, want %s, got %s`, want, result)
	}
}

func TestInsertSlice(t *testing.T) {
	columns := []string{"Name", "Id"}
	values := []interface{}{1, "3"}
	stmt, args, err := Insert("example").Columns(columns...).Values(values...).ToSql()
	if err != nil {
		t.Fatalf(`Insert("example").Columns("Name", "Id").Values(1, "3").ToSql() - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := fmt.Sprintf("INSERT INTO example (Name,Id) (?,'?') [1 3]")

	if result != want {
		t.Fatalf(`Insert("example").Columns("Name", "Id").Values(1, "3").ToSql() failed, want %s, got %s`, want, result)
	}
}
