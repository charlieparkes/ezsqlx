package ezsqlx

import (
	"testing"

	"github.com/charlieparkes/go-structs"
	"github.com/stretchr/testify/assert"
)

func TestColumns(t *testing.T) {
	type testModel struct {
		Foo     int    `db:"foo" constraint:"pk"`
		Bar     string `db:"bar"`
		FooBar  string `db:"wiz_bang"`
		NotHere string
	}

	assert.Equal(t, []string{"foo", "bar", "wiz_bang", "not_here"}, Columns(structs.Fields(&testModel{})))
}

func TestPrimaryKey(t *testing.T) {
	type testModel struct {
		Foo     int    `db:"foo" constraint:"pk"`
		Bar     string `db:"bar"`
		FooBar  string `db:"wiz_bang"`
		NotHere string
	}

	assert.Equal(t, []string{"foo"}, primaryKey(structs.Fields(&testModel{})))
}
