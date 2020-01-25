package format

import (
	"testing"

	"gotest.tools/assert"
)

func Test_ArgsToString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		args []interface{}
		s    string
	}{
		"no args":        {nil, ""},
		"one arg":        {[]interface{}{2}, "2"},
		"two args":       {[]interface{}{2, 3}, "2 3"},
		"format and arg": {[]interface{}{"one is %d", 1}, "one is 1"},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			s := ArgsToString(tc.args...)
			assert.Equal(t, tc.s, s)
		})
	}
}
