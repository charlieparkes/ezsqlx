package ezsqlx

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlieparkes/ezsqlx/test"
)

func TestInsert(t *testing.T) {
	var err error

	fixture := test.PostgresDatabase{}
	err = fixture.SetUp()
	assert.Nil(t, err)
	defer fixture.TearDown()

	db, err := fixture.Settings.Connect()
	assert.Nil(t, err)
	defer db.Close()

	test.CreateDummyTable(db)

	newRow := &test.FooBar{Message: "confused unga bunga"}
	rows, err := Insert(db, "foobar", newRow, []string{"id", "created"})
	assert.Nil(t, err)

	foobar := &test.FooBar{}
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
