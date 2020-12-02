# ezsqlx

**ezsqlx** is a library which provides helpers for `jmoiron/sqlx` which simplify basic database actions. **ezsqlx** will improve your experience without incurring the penalties of an ORM.

##### Features

* Insert
* Update
* Model Fields

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
> ```go
> Insert(
>   db *sqlx.DB,
>   table string,
>   model interface{},
>   excludedFields []string
> )(*sqlx.Rows, error)
> ```

###### Example
```go
newRow := &FooBar{Message: "confused unga bunga"}
rows, err := Insert(db, "foobar", newRow, []string{"id", "created"})
```

### Update

> ```go
> Update(
>  	db *sqlx.DB,
>  	table string,
>  	model interface{},
>  	where string,
>  	excludedFields []string
> ) (sql.Result, error) 
> ```

###### Example
```go
newRow := &FooBar{Id: 1, Message: "confused unga bunga"}
_, err = Update(db, "foobar", foobar, fmt.Sprintf("id=%v",foobar.Id), []string{"id", "created"})
```

### Fields

> ```go
> Fields(
>   values interface{}
> ) []string
> ```

###### Example

```go
fields := Fields(model)
```

