package ezsqlx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyModel struct {
	Foo     int    `db:"foo"`
	Bar     string `db:"bar"`
	FooBar  string `db:"wiz_bang"`
	NotHere string
}

func TestFields(t *testing.T) {
	dm := DummyModel{}
	expectedFields := []string{"foo", "bar", "wiz_bang"}
	fields := Fields(dm)
	for i, f := range fields {
		assert.Equal(t, f, expectedFields[i], "field names should be equal")
	}
}
