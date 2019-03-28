# go-rules
Simple rules engine

## Examples


```go
// (a == 1 && b == 2) || (c == 3 && d == 4)
r := []byte(`{ "comparator": "||", "rules": [ { "comparator": "&&", "rules": [ { "var": "a", "op": "==", "val": 1 }, { "var": "b", "op": "==", "val": 2 } ] }, { "comparator": "&&", "rules": [ { "var": "c", "op": "==", "val": 3 }, { "var": "d", "op": "==", "val": 4 } ] } ] }`)

// parse from JSON
expr, err := rules.ParseFromJSON(r)
if err != nil {
    // ...
}

params := map[string]interface{}{
    "a": 1,
    "b": 2,
    "c": 0,
    "d": 0,
}

// passing data to evaluate expression
res, err := rules.Evaluate(expr, params)
if err != nil {
    // ...
}

// res is true
fmt.Println(res)
```