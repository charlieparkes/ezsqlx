package ezsqlx

import (
	"log"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/tovala/ezsqlx/test"
)

type FooBar struct {
	Id      int        `db:"id"`
	Message string     `db:"message"`
	Flip    bool       `db:"flip"`
	Created *time.Time `db:"created"`
}

func createTable(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec(`CREATE TABLE foobar (
		id SERIAL PRIMARY KEY,
		message TEXT,
		flip BOOL NOT NULL DEFAULT false,
		created TIMESTAMP DEFAULT now()
	)`)
	tx.Commit()
}

func TestInsert(t *testing.T) {
	var err error

	fixture := test.PostgresDatabase{}
	err = fixture.SetUp()
	assert.Nil(t, err)
	defer fixture.TearDown()

	db, err := fixture.Settings.Connect()
	assert.Nil(t, err)
	defer db.Close()

	createTable(db)

	newRow := &FooBar{Message: "confused unga bunga"}
	rows, err := Insert(db, "foobar", newRow, []string{"id"})
	assert.Nil(t, err)

	foobar := &FooBar{}
	for rows.Rows.Next() {
		err = rows.StructScan(foobar)
		if err != nil {
			log.Fatalln(err)
		}
	}

	assert.Equal(t, 1, foobar.Id)
	assert.Equal(t, "confused unga bunga", foobar.Message)
	assert.Equal(t, false, foobar.Flip)
}

// func TestInsertMany(t *testing.T) {
// 	var err error

// 	fixture := test.PostgresDatabase{}
// 	err = fixture.SetUp()
// 	assert.Nil(t, err)
// 	defer fixture.TearDown()

// 	db, err := fixture.Settings.Connect()
// 	assert.Nil(t, err)
// 	defer db.Close()

// 	createTable(db)

// 	newRows := []FooBar{
// 		FooBar{Message: "confused unga bunga"},
// 		FooBar{Message: "the welfare of humanity is always the alibi of tyrants"},
// 	}

// 	query := InsertManyQuery(db, "foobar", FooBar{}, []string{"id"})
// 	rows, err := db.NamedQuery(query, newRows)
// 	assert.Nil(t, err)

// 	scannedRows := []FooBar{}
// 	for rows.Rows.Next() {
// 		foobar := FooBar{}
// 		err = rows.StructScan(foobar)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		scannedRows = append(scannedRows, foobar)
// 	}

// 	assert.Equal(t, 1, scannedRows[0].Id)
// 	assert.Equal(t, "confused unga bunga", scannedRows[0].Message)
// 	assert.Equal(t, false, scannedRows[0].Flip)
// }
