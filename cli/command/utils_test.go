package command

import (
	"testing"

	"gotest.tools/assert"
)

func TestReplaceArgs(t *testing.T) {
	out, ok := ReplaceArgs([]string{"foo", "bar"}, []string{"baz"}, []string{"abc", "foo", "bar", "bax"}, -1)
	assert.Assert(t, ok)
	assert.DeepEqual(t, []string{"abc", "baz", "bax"}, out)

	out, ok = ReplaceArgs([]string{"foo", "bar"}, []string{"baz"}, []string{"foo"}, -1)
	assert.Assert(t, !ok)
	assert.DeepEqual(t, []string{"foo"}, out)

	out, ok = ReplaceArgs([]string{"foo", "bar"}, []string{"baz"}, []string{"abc", "foo", "bar", "bax"}, 0)
	assert.Assert(t, !ok)
	assert.DeepEqual(t, []string{"abc", "foo", "bar", "bax"}, out)

	out, ok = ReplaceArgs([]string{"foo", "bar"}, []string{"baz"}, []string{"foo", "bar", "bax"}, 0)
	assert.Assert(t, ok)
	assert.DeepEqual(t, []string{"baz", "bax"}, out)

	out, ok = ReplaceArgs([]string{"foo", "bar"}, nil, []string{"abc", "foo", "bar", "baz"}, -1)
	assert.Assert(t, ok)
	assert.DeepEqual(t, []string{"abc", "baz"}, out)

	out, ok = ReplaceArgs(nil, []string{"baz"}, []string{"foo"}, -1)
	assert.Assert(t, !ok)
	assert.DeepEqual(t, []string{"foo"}, out)
}
