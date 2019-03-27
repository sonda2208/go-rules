package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rules "github.com/sonda2208/go-rules"
)

func TestParser(t *testing.T) {
	tests := []struct {
		rules   string
		isError bool
	}{
		{
			`{ "comparator": "and", "rules": [ { "comparator": "or", "rules": [ { "var": "a", "op": "eq", "val": "xavier" }, { "var": "b", "op": "eq", "val": 1 } ] }, { "var": "c", "op": "eq", "val": true } ] }`,
			false,
		},
		{
			`{ "comparator": "or", "rules": [ { "comparator": "and", "rules": [ { "var": "a", "op": "eq", "val": 1 }, { "var": "b", "op": "eq", "val": 2 } ] }, { "comparator": "and", "rules": [ { "var": "c", "op": "eq", "val": 3 }, { "var": "d", "op": "eq", "val": 4 } ] } ] }`,
			false,
		},
		{
			`{ "comparator": "or", "rules": [ { "comparator": "and", "rules": [ { "var": "a", "op": "eq", "val": 1 }, { "var": "b", "op": "eq", "val": 2 } ] }, { "comparator": "and", "rules": [ { "var": "c", "op": "eq", "val": 3 }, { "var": "d", "op": "eq", "val": 4 } ] } ], "var": "a", }`,
			true,
		},
		{
			`{ "comparator": "or", "rules": [ { "comparator": "and", "rules": [ { "var": "a", "op": "eq", "val": 1 }, { "var": "b", "op": "eq", "val": 2 } ] }, { "comparator": "and", "rules": [ { "var": "c", "op": "eq", "val": 3 }, { "var": "d", "op": "eq", "val": 4 } ] } ], "op": "?" }`,
			true,
		},
		{
			`{ "comparator": "or", "rules": [ { "comparator": "and", "rules": [ { "var": "a", "op": "eq", "val": 1 }, { "var": "b", "op": "eq", "val": 2 } ] }, { "comparator": "and", "rules": [ { "var": "c", "op": "eq", "val": 3 }, { "var": "d", "op": "eq", "val": 4 } ] } ], "val": 1 }`,
			true,
		},
		{
			`{ "comparator": "or", "rules": [ { "comparator": "and", "rules": [ { "var": "a", "op": "eq", "val": 1 }, { "var": "b", "op": "eq", "val": 2 } ] }, { "comparator": "and", "rules": [ { "var": "c", "op": "eq", "val": 3 }, { "var": "d", "op": "eq", "val": 4 } ] } ], "var": "a", "op": "eq", "val": 1 }`,
			true,
		},
	}

	for _, test := range tests {
		expr, err := rules.ParseFromJSON([]byte(test.rules))
		if test.isError {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.NotNil(t, expr)
		}
	}
}
