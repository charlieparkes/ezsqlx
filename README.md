# ezsqlx

**ezsqlx** is a library which provides helpers for `jmoiron/sqlx` which simplify basic database actions. **ezsqlx** will improve your experience without incurring the penalties of an ORM.

##### Features

* Insert
* Update
* Connection management
* Get database fields from a struct

## Examples

Given a basic model struct, **ezsqlx** simplifies basic operations.

```go
type FooBar struct {
    Id      int        `db:"id"`
    Message string     `db:"message"`
    Flip    bool       `db:"flip"`
    Created *time.Time `db:"created"`
}
```

### Insert

###### Interface
```go
Insert(
    db *sqlx.DB,
    table string,
    model interface{},
    excludedFields []string
) (*sqlx.Rows, error)
```

###### Example
```go
newRow := &FooBar{Message: "confused unga bunga"}
rows, err := Insert(db, "foobar", newRow, []string{"id", "created"})
```

### Update

###### Interface
```go
Update(
    db *sqlx.DB,
    table string,
    model interface{},
    where string,
    excludedFields []string
) (sql.Result, error) 
```

###### Example
```go
updatedRow := &FooBar{Id: 1, Message: "pc master race", Flip: true}
where := fmt.Sprintf("id=%v", updatedRow.Id)
_, err = Update(db, "foobar", updatedRow, where, []string{"id", "created"})
```

### Fields

###### Interface
```go
Fields(
    values interface{}
) []string
```

###### Example
```go
fields := Fields(model)
```

### Connections

`ezsqlx.ConnectionSettings` abstracts away basic Postgres connection operations. Check `connections.go` for a full list of helpers.

###### Example
```go
cs := &ConnectionSettings{
    Host: "localhost",
    Port: "1234",
    User: "postgres",
    Password: "postgres",
    Database: "my_database"
}

db := cs.Open()
defer db.Close()
```