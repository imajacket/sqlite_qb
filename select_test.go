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
	want := "SELECT * FROM example []"

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
	want := "SELECT * FROM example WHERE 'id' = '?' [test_id]"

	if result != want {
		t.Fatalf(`Select("*").From("example").Where(Compare(Eq, map[string]interface{}{"id": "test_id",})).ToSQL() failed, want %s, got %s`, want, result)
	}
}

func TestSelectLike(t *testing.T) {
	stmt, args, err := Select("*").From("example").Where(Like(map[string]string{
		"name": "ington",
	})).ToSQL()
	if err != nil {
		t.Fatalf(`Select("*").From("example").Where(Like(map[string]string{"name": "ton",})).ToSQL() failed - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := "SELECT * FROM example WHERE 'name' LIKE '?' [%ington%]"

	if result != want {
		t.Fatalf(`Select("*").From("example").Where(Like(map[string]string{"name": "ton",})).ToSQL(), want %s, got %s`, want, result)
	}
}

func TestSelectBetween(t *testing.T) {
	stmt, args, err := Select("*").From("example").Where(Between(map[string][]interface{}{
		"age": {20, 29},
	})).ToSQL()
	if err != nil {
		t.Fatalf(`Select("*").From("example").Where(Between(map[string][]interface{}{"age": {20, 29},})).ToSQL() failed - %v`, err)
	}

	result := fmt.Sprintf("%v %v", stmt, args)
	want := "SELECT * FROM example WHERE 'age' BETWEEN ? and ? [20 29]"

	if result != want {
		t.Fatalf(`Select("*").From("example").Where(Between(map[string][]interface{}{"age": {20, 29},})).ToSQL(), want %s, got %s`, want, result)
	}
}
