package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rules "github.com/sonda2208/go-rules"
)

func TestJSONParser(t *testing.T) {
	testcases := []struct {
		rules   string
		isError bool
	}{
		{
			`{ "var": "strlen(a)", "op": ">", "val": 1 }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": false }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": 0 }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": -1 }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": 1 }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": "0" }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": "2h" }`,
			false,
		},
		{
			`{ "var": "a", "op": "==", "val": "2006-01-02T15:04:05Z" }`,
			false,
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
		{
			`{ "comparator": "&&", "rules": [ { "comparator": "||", "rules": [ { "var": "a", "op": "==", "val": "xavier" }, { "var": "b", "op": "==", "val": 1 } ] }, { "var": "c", "op": "==", "val": true } ] }`,
			false,
		},
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ] }`,
			false,
		},
		{
			`{ "var": "a", "op": "contains", "val": 1 }`,
			false,
		},
		{
			`{ "var": "a", "op": "contains", "val": "abc" }`,
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
			`{ "var": "a", "op": "=", "val": 1 }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [] }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [0, "1", 2] }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [true, false] }`,
			true,
		},
		{
			`{ "var": "a", "op": "in", "val": [0, 1, false] }`,
			true,
		},
		{
			`{ "var": "a", "op": "", "val": "2h" }`,
			true,
		},
		{
			`{ "var": "", "op": "==", "val": "2h" }`,
			true,
		},
		{
			`{ "var": "a", "op": "=="}`,
			true,
		},
		{
			`{"comparator": "||" }`,
			true,
		},
		{
			`{ "comparator": "||", "rules": [] }`,
			true,
		},
		{
			`{ "comparator": "&&", "rules": [ { "comparator": "||", "rules": [ { "var": "", "op": "==", "val": "xavier" }, { "var": "b", "op": "==", "val": 1 } ] }, { "var": "c", "op": "in", "val": [2, 3, 4] } ] }`,
			true,
		},
		{
			`{ "var": "strlen(a", "op": ">", "val": 1 }`,
			true,
		},
		{
			`{ "var": "a", "op": "contain", "val": "abc" }`,
			true,
		},
	}

	for _, tc := range testcases {
		expr, err := rules.ParseFromJSON([]byte(tc.rules))
		if tc.isError {
			assert.Error(t, err, tc.rules)
		} else {
			require.NoError(t, err)
			assert.NotNil(t, expr, tc.rules)
		}
	}
}
