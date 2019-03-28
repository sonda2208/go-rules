package rules_test

import (
	"testing"
	"time"

	rules "github.com/sonda2208/go-rules"
)

func BenchmarkComplexExpression(bench *testing.B) {
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
			  { "var": "d", "op": ">=", "val": 4 }
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

	expr, _ := rules.ParseFromJSON([]byte(expression))

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		rules.Evaluate(expr, parameters)
	}
}
