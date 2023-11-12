# SQLite Query Builder

This is very rudimentary, only writing this to practice writing Go and working with SQL.


```go
// Select
stmt, args, err := qb.Select("*").From("example").ToSQL()

fmt.Printf("%v %v", stmt, args)
// SELECT * FROM example []

stmt, args, err := qb.Select("*").From("example").Where(qb.Compare(qb.Eq, map[string]interface{}{
    "id": "test_id",
})).ToSQL()
// SELECT * FROM example WHERE 'id' = '?' [test_id]

stmt, args, err := Select("*").From("example").Where(Like(map[string]string{
    "name": "ington",
})).ToSQL()

fmt.Printf("%v %v", stmt, args) 
// SELECT * FROM example WHERE 'name' LIKE '?' [%ington%]

stmt, args, err := Select("*").From("example").Where(Between(map[string][]interface{}{
    "age": {20, 29},
})).ToSQL()

fmt.Printf("%v %v", stmt, args)
// SELECT * FROM example WHERE 'age' BETWEEN ? and ? [20 29]

// Insert
stmt, args, err := qb.Insert("example").Columns("Name", "Id").Values(1, "3").ToSql()

fmt.Printf("%v %v", stmt, args)
// INSERT INTO example (Name,Id) (?,'?') [1 3]
```