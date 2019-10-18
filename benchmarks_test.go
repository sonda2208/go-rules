package rules_test

import (
	"testing"
	"time"

	rules "github.com/sonda2208/go-rules"
	"github.com/stretchr/testify/require"
)

func BenchmarkEvaluateFromJSON(b *testing.B) {
	expression := `
	{
		"comparator": "||",
		"rules": [
		  {
			"comparator": "&&",
			"rules": [
			  { "var": "a", "op": "<", "val": "2019-03-28T11:39:43+07:00" },
			  { "var": "b", "op": "in", "val": [1, 2, 3] }
			]
		  },
		  {
			"comparator": "&&",
			"rules": [
			  { "var": "c", "op": "!=", "val": "string" },
			  { "var": "d", "op": ">=", "val": 4 },
			  { "var": "strlen(c)", "op": ">", "val": 0 }
			]
		  }
		]
	}
	`

	parameters := map[string]interface{}{
		"a": time.Now(),
		"b": 1,
		"c": "number",
		"d": 5,
	}

	functions := map[string]rules.Function{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
	}

	expr, err := rules.ParseFromJSON([]byte(expression))
	require.NoError(b, err)
	require.NotNil(b, expr)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rules.Evaluate(expr, parameters, functions)
	}
}

func BenchmarkEvaluateFromString(b *testing.B) {
	input := `a == "2006-01-02T15:04:05Z" && (b > 0 || c == "string") && d < 100 && strlen(c) > 0`
	expr, err := rules.ParseFromString(input)
	require.NoError(b, err)
	require.NotNil(b, expr)

	parameters := map[string]interface{}{
		"a": time.Now(),
		"b": 1,
		"c": "number",
		"d": 5,
	}

	functions := map[string]rules.Function{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rules.Evaluate(expr, parameters, functions)
	}
}
