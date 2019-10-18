package rules_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/icrowley/fake"

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
	funcs       map[string]rules.Function
}

func defaultTime() time.Time {
	dt, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return dt
}

func getDuration(dur string) time.Duration {
	res, _ := time.ParseDuration(dur)
	return res
}

func getTime(t string) time.Time {
	dt, _ := time.Parse(time.RFC3339, t)
	return dt
}

func TestEvaluate(t *testing.T) {
	runTestCase := func(testcases []TestCase) {
		for _, tc := range testcases {
			expr, err := rules.ParseFromString(tc.rules)
			require.NoError(t, err)
			require.NotNil(t, expr)

			for i, e := range tc.evaluations {
				res, err := rules.Evaluate(expr, e.params, tc.funcs)
				if e.isError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err, fmt.Sprintf("%v %d:%v", tc.rules, i, e.params))
					assert.Equal(t, e.expected, res, fmt.Sprintf("%v %d:%v", tc.rules, i, e.params))
				}
			}
		}
	}

	t.Run("General", func(t *testing.T) {
		t.Run("Nil expression", func(t *testing.T) {
			_, err := rules.Evaluate(nil, nil, nil)
			assert.Error(t, err)
		})

		t.Run("Invalid returned type", func(t *testing.T) {
			input := `strlen("???")`
			expr, err := rules.ParseFromString(input)
			require.NoError(t, err)
			require.NotNil(t, expr)

			_, err = rules.Evaluate(expr, nil, map[string]rules.Function{
				"strlen": func(args ...interface{}) (interface{}, error) {
					length := len(args[0].(string))
					return (float64)(length), nil
				},
			})
			assert.Error(t, err)
		})

		t.Run("Parameter not found", func(t *testing.T) {
			input := `a == 1`
			expr, err := rules.ParseFromString(input)
			require.NoError(t, err)
			require.NotNil(t, expr)

			_, err = rules.Evaluate(expr, nil, nil)
			assert.Error(t, err)
		})

		t.Run("Custom function not found", func(t *testing.T) {
			input := `strlen("???")`
			expr, err := rules.ParseFromString(input)
			require.NoError(t, err)
			require.NotNil(t, expr)

			_, err = rules.Evaluate(expr, nil, nil)
			assert.Error(t, err)
		})

		t.Run("Custom function returns error", func(t *testing.T) {
			input := `strlen("???")`
			expr, err := rules.ParseFromString(input)
			require.NoError(t, err)
			require.NotNil(t, expr)

			_, err = rules.Evaluate(expr, nil, map[string]rules.Function{
				"strlen": func(args ...interface{}) (interface{}, error) {
					return nil, errors.New("unknown")
				},
			})
			assert.Error(t, err)
		})
	})

	t.Run("Parse from JSON with custom function", func(t *testing.T) {
		input := `{ "var": "strlen(a)", "op": ">", "val": 1 }`
		expr, err := rules.ParseFromJSON(json.RawMessage(input))
		require.NoError(t, err)
		require.NotNil(t, expr)

		res, err := rules.Evaluate(
			expr,
			map[string]interface{}{
				"a": "???",
			},
			map[string]rules.Function{
				"strlen": func(args ...interface{}) (interface{}, error) {
					length := len(args[0].(string))
					return (float64)(length), nil
				},
			})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("Parse from JSON with nested custom functions", func(t *testing.T) {
		input := `{ "var": "strlen(randomChars(n))", "op": ">", "val": 0 }`
		expr, err := rules.ParseFromJSON(json.RawMessage(input))
		require.NoError(t, err)
		require.NotNil(t, expr)

		res, err := rules.Evaluate(
			expr,
			map[string]interface{}{
				"n": 3,
			},
			map[string]rules.Function{
				"strlen": func(args ...interface{}) (interface{}, error) {
					length := len(args[0].(string))
					return (float64)(length), nil
				},
				"randomChars": func(args ...interface{}) (interface{}, error) {
					n := args[0].(float64)
					return fake.CharactersN(int(n)), nil
				},
			})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("Convert literal", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a < 2`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": int32(1),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": int64(1),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": float32(1),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": float64(1),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 2,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []int{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []int32{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []int64{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []float32{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []float64{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []string{"0", "1", "2"},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{0, 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{int32(0), 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{int64(0), 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{float32(0), 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{float64(0), 1, 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{0, "1", 2},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{"0", "1", "2"},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []interface{}{"0", 1, "2"},
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": struct{}{},
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("OR operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a >= 1 || b < 2`,
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
				nil,
			},
			{
				`false || true`,
				[]Evaluation{
					{
						nil,
						true,
						false,
					},
				},
				nil,
			},
			{
				`1 || true`,
				[]Evaluation{
					{
						nil,
						false,
						true,
					},
				},
				nil,
			},
			{
				`1 || a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "0",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`false || a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": true,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("AND operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`false && 1`,
				[]Evaluation{
					{
						nil,
						false,
						true,
					},
				},
				nil,
			},
			{
				`true && true`,
				[]Evaluation{
					{
						nil,
						true,
						false,
					},
				},
				nil,
			},
			{
				`false && true`,
				[]Evaluation{
					{
						nil,
						false,
						false,
					},
				},
				nil,
			},
			{
				`true && a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": true,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a == 1 && b == 2 && c == 3`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
							"b": "2",
							"c": 3,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
							"b": "",
							"c": 3,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
							"b": 2,
							"c": 3,
						},
						true,
						false,
					},
				},
				nil,
			},
			{
				`a > 1 && b < 2`,
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
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("EQ operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a == "0"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "0",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "1",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a == "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "",
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" == a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2006-01-02T15:04:05Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`true == true`,
				[]Evaluation{
					{
						nil,
						true,
						false,
					},
				},
				nil,
			},
			{
				`true == false`,
				[]Evaluation{
					{
						nil,
						false,
						false,
					},
				},
				nil,
			},
			{
				`a == true`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": true,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`true == a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": true,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a == 1h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`1s == a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "1s",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2s",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2s"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("NEQ operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a != 1`,
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
				nil,
			},
			{
				`a != "0"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "0",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "1",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a != 2h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a != false`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": true,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "1",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a != "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "",
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" != a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2006-01-02T15:04:01Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`1s != a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2s"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "1s",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2s",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": true,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("LT operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a < 2`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 2,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a < "abc"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "ABC",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "abc",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a < 2h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						false,
						false,
					},
				},
				nil,
			},
			{
				`2h < a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "3h",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": []int{0, 1, 2},
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a < "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" < a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now().Format(time.RFC3339),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("LTE operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a <= 2`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 2,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a <= "abc"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "ABC",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "abc",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a <= 2h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`2h <= a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "3h",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a <= "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2006-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" <= a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now().Format(time.RFC3339),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2006-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2006-01-02T15:04:05Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("GT operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a > 2`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 2,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 3,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a > "abc"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "d",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "ABC",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "abc",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a > 2h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						true,
						false,
					},
				},
				nil,
			},
			{
				`2h > a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "3h",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2h",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a > "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" > a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now().Format(time.RFC3339),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("GTE operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a >= 2`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 2,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 3,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a >= "abc"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "d",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "ABC",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "abc",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a >= 2h`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						true,
						false,
					},
				},
				nil,
			},
			{
				`2h >= a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getDuration("3h"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("2h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getDuration("1h"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "3h",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": "2h",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "1h",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`a >= "2006-01-02T15:04:05Z"`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2006-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 1,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
			{
				`"2006-01-02T15:04:05Z" >= a`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": time.Now(),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": time.Now().Format(time.RFC3339),
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2005-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": getTime("2006-01-02T15:04:05Z"),
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2005-01-02T15:04:05Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "2006-01-02T15:04:05Z",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("IN operator", func(t *testing.T) {
		testcases := []TestCase{
			{
				`a in [0, 1, 2]`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 4,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": false,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": "?",
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				map[string]rules.Function{},
			},
			{
				`a in ["0", "1", "2"]`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "1",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "4",
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 0,
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": defaultTime(),
						},
						false,
						true,
					},
					{
						map[string]interface{}{
							"a": getDuration("1s"),
						},
						false,
						true,
					},
				},
				map[string]rules.Function{},
			},
		}
		runTestCase(testcases)
	})

	t.Run("Parentheses", func(t *testing.T) {
		testcases := []TestCase{
			{
				`(a == 1 && b == 2) || (c == 3 && d == 4)`,
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
				nil,
			},
			{
				`a == 1 && (b == 2 || c == 3)`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": 1,
							"b": 2,
							"c": 3,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 1,
							"b": 1,
							"c": 3,
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": 0,
							"b": 2,
							"c": 3,
						},
						false,
						false,
					},
				},
				nil,
			},
		}
		runTestCase(testcases)
	})

	t.Run("Custom function", func(t *testing.T) {
		testcases := []TestCase{
			{
				`strlen(a) > 0 && isCalled()`,
				[]Evaluation{
					{
						map[string]interface{}{
							"a": "???",
						},
						true,
						false,
					},
					{
						map[string]interface{}{
							"a": "",
						},
						false,
						false,
					},
				},
				map[string]rules.Function{
					"strlen": func(args ...interface{}) (interface{}, error) {
						length := len(args[0].(string))
						return (float64)(length), nil
					},
					"isCalled": func(args ...interface{}) (interface{}, error) {
						return true, nil
					},
				},
			},
			{
				`sum(a, b) > 0`,
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
							"b": 0,
						},
						false,
						false,
					},
					{
						map[string]interface{}{
							"a": 0,
							"b": "?",
						},
						false,
						true,
					},
					{
						nil,
						false,
						true,
					},
				},
				map[string]rules.Function{
					"sum": func(args ...interface{}) (interface{}, error) {
						if len(args) != 2 {
							return nil, errors.New("invalid arguments")
						}

						a, ok := args[0].(float64)
						if !ok {
							return nil, errors.New("invalid arguments")
						}

						b, ok := args[1].(float64)
						if !ok {
							return nil, errors.New("invalid arguments")
						}

						return a + b, nil
					},
				},
			},
		}
		runTestCase(testcases)
	})
}

func TestExamples(t *testing.T) {
	t.Run("From JSON", func(t *testing.T) {
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
		require.NoError(t, err)

		res, err := rules.Evaluate(expr, parameters, functions)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("From string", func(t *testing.T) {
		input := `a == "2006-01-02T15:04:05Z" && (b > 0 || c == "string") && d < 100 && strlen(c) > 0`

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

		expr, err := rules.ParseFromString(input)
		require.NoError(t, err)
		require.NotNil(t, expr)

		res, err := rules.Evaluate(expr, parameters, functions)
		require.NoError(t, err)
		assert.False(t, res)
	})
}
