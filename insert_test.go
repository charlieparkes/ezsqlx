package ezsqlx

import (
	"log"
	"testing"

	"github.com/charlieparkes/ezsqlx/test"
	"github.com/stretchr/testify/assert"
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

func TestInsertQuery(t *testing.T) {
	type testModel struct {
		Foo     string `constraint:"pk"`
		Bar     int    `db:"barry"`
		WizBang bool
	}

	sql := InsertQuery("my_table", testModel{})
	assert.Equal(t, "INSERT INTO my_table (\"foo\", \"barry\", \"wiz_bang\") VALUES (:foo, :barry, :wiz_bang)", sql)
}

func TestUpsertQuery(t *testing.T) {
	type testModel struct {
		Foo     string `constraint:"pk"`
		Bar     int    `db:"barry"`
		WizBang bool
	}

	sql := UpsertQuery("my_table", testModel{}, "my_table_pk")
	assert.Equal(t, "INSERT INTO my_table (\"foo\", \"barry\", \"wiz_bang\") VALUES (:foo, :barry, :wiz_bang) ON CONFLICT ON CONSTRAINT my_table_pk DO UPDATE SET barry = EXCLUDED.barry, wiz_bang = EXCLUDED.wiz_bang", sql)
}
