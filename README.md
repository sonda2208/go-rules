# go-rules
Simple rules engine

## Examples


```go
// (a < '2019-03-28T11:39:43+07:00' && b in (1, 2, 3)) || (c != "string" && d >= 4)
r := []byte(`
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
`)

// parse from JSON
expr, err := rules.ParseFromJSON(r)
if err != nil {
    // ...
}

params := map[string]interface{}{
    "a": time.Now(),
    "b": 1,
    "c": "number",
    "d": 5,
}

// passing data to evaluate expression
res, err := rules.Evaluate(expr, params)
if err != nil {
    // ...
}

// res is true
fmt.Println(res)
```