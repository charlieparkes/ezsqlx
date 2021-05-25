package ezsqlx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "foo_bar", snakeCase("FooBar"))
	assert.Equal(t, "foo_bar", snakeCase("fooBar"))
}

func TestRemove(t *testing.T) {
	x := []string{"a", "b", "c", "d", "e"}

	x = remove(x, "b")
	assert.Equal(t, []string{"a", "c", "d", "e"}, x)

	x = remove(x, "e")
	assert.Equal(t, []string{"a", "c", "d"}, x)

	x = remove(x, "a")
	assert.Equal(t, []string{"c", "d"}, x)
}
