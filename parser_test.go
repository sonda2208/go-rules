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
			`{ "comparator": "&&", "rules": [ { "comparator": "||", "rules": [ { "var": "a", "op": "==", "val": "xavier" }, { "var": "b", "op": "==", "val": 1 } ] }, { "var": "c", "op": "==", "val": true } ] }`,
			false,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ] }`,
			false,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ], "var": "a", }`,
			true,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ], "op": "?" }`,
			true,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ], "val": 1 }`,
			true,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ], "var": "a", "op": "==", "val": 1 }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [] }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [1, 2, 3] }`,
			false,
		},
		{
			`{ "var": "a", "op": "in", "val": ["1", "2", "3"] }`,
			false,
		},
		{
			`{ "comparator": "&&", "rules": [ { "comparator": "||", "rules": [ { "var": "a", "op": "==", "val": "xavier" }, { "var": "b", "op": "==", "val": 1 } ] }, { "var": "c", "op": "in", "val": [2, 3, 4] } ] }`,
			false,
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
