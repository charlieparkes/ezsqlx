package ezsqlx

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlieparkes/ezsqlx/test"
)

func TestUpdate(t *testing.T) {
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
	rows, err := Insert(db, "foobar", newRow, []string{"id"})
	assert.Nil(t, err)

	foobar := &test.FooBar{}
	for rows.Rows.Next() {
		err = rows.StructScan(foobar)
		if err != nil {
			log.Fatalln(err)
		}
	}

	foobar.Message = "pc master race"
	foobar.Flip = true

	_, err = Update(db, "foobar", foobar, fmt.Sprintf("id=%v", foobar.Id), []string{"id", "created"})
	assert.Nil(t, err)

	updatedFoobar := &test.FooBar{}
	err = db.Get(updatedFoobar, "select * from foobar where id=$1", foobar.Id)
	assert.Nil(t, err)
	assert.Equal(t, true, updatedFoobar.Flip)
}
