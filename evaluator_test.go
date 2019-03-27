package rules_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rules "github.com/sonda2208/go-rules"
)

type Evaluation struct {
	params   map[string]interface{}
	expected bool
	isError  bool
}

type TestCase struct {
	rules       string
	evaluations []Evaluation
}

func TestEvaluator(t *testing.T) {
	tests := []TestCase{
		{
			`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ] }`,
			[]Evaluation{
				{
					map[string]interface{}{
						"a": 1,
						"b": 2,
						"c": 0,
						"d": 0,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 0,
						"c": 3,
						"d": 4,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": 2,
						"c": 3,
						"d": 4,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 2,
						"c": 0,
						"d": 0,
					},
					false,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": 0,
						"c": 3,
						"d": 0,
					},
					false,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 0,
						"c": 0,
						"d": 0,
					},
					false,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": 2,
						"c": 0,
						"d": "string",
					},
					false,
					true,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": 2,
						"c": true,
						"d": 0,
					},
					false,
					true,
				},
			},
		},
		{
			`{ "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": "2" }, { "var": "c", "op": "==", "val": 3 } ] }`,
			[]Evaluation{
				{
					map[string]interface{}{
						"a": 1,
						"b": "2",
						"c": 3,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": "",
						"c": 3,
					},
					false,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
						"b": 2,
						"c": 3,
					},
					false,
					true,
				},
			},
		},
		{
			`{ "comparator": "&&", "rules": [ { "var": "a", "op": ">", "val": 1 }, { "var": "b", "op": "<", "val": 2 } ] }`,
			[]Evaluation{
				{
					map[string]interface{}{
						"a": 2,
						"b": 1,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 1,
					},
					false,
					false,
				},
				{
					map[string]interface{}{
						"a": 2,
						"b": -1,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 2,
						"b": 2,
					},
					false,
					false,
				},
			},
		},
		{
			`{ "comparator": "||", "rules": [ { "var": "a", "op": ">=", "val": 1 }, { "var": "b", "op": "<", "val": 2 } ] }`,
			[]Evaluation{
				{
					map[string]interface{}{
						"a": 1,
						"b": 1,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 1,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 0,
						"b": 2,
					},
					false,
					false,
				},
			},
		},
		{
			`{ "var": "a", "op": "!=", "val": 1 }`,
			[]Evaluation{
				{
					map[string]interface{}{
						"a": 2,
					},
					true,
					false,
				},
				{
					map[string]interface{}{
						"a": 1,
					},
					false,
					false,
				},
			},
		},
	}

	for _, test := range tests {
		expr, err := rules.ParseFromJSON([]byte(test.rules))
		require.NoError(t, err)
		require.NotNil(t, expr)

		for _, e := range test.evaluations {
			res, err := rules.Evaluate(expr, e.params)
			if e.isError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				if !assert.Equal(t, e.expected, res) {
					t.Log(expr.String())
				}
			}
		}
	}
}
