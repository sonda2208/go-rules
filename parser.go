package rules

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

type Rule struct {
	Comparator Op          `json:"comparator,omitempty"`
	Rules      []Rule      `json:"rules,omitempty"`
	Var        string      `json:"var,omitempty"`
	Op         Op          `json:"op,omitempty"`
	Val        interface{} `json:"val,omitempty"`
}

func (r *Rule) IsValid() error {
	if len(r.Comparator) > 0 {
		if len(r.Rules) == 0 {
			return errors.New(`"Rules" must not be empty in complex mode`)
		}

		if len(r.Var) > 0 {
			return errors.New(`"Var" must be empty in complex mode`)
		}

		if len(r.Op) > 0 {
			return errors.New(`"Op" must be empty in complex mode`)
		}

		if r.Val != nil {
			return errors.New(`"Val" must be nil in complex mode`)
		}
	} else {
		if len(r.Var) == 0 {
			return errors.New(`"Var" must not be empty`)
		}

		if len(r.Op) == 0 {
			return errors.New(`"Op" must not be empty`)
		}

		if r.Val == nil {
			return errors.New(`"Val" must not be nil`)
		}
	}

	return nil
}

func ParseFromJSON(raw json.RawMessage) (Expr, error) {
	rule := Rule{}
	err := json.Unmarshal(raw, &rule)
	if err != nil {
		return nil, err
	}

	expr, err := parseExpr(&rule)
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func parseExpr(rule *Rule) (Expr, error) {
	err := rule.IsValid()
	if err != nil {
		return nil, err
	}

	if len(rule.Rules) > 0 {
		var res Expr
		for _, r := range rule.Rules {
			e, err := parseExpr(&r)
			if err != nil {
				return nil, err
			}

			if res == nil {
				res = e
			} else {
				res = &BinaryExpr{
					Op:  rule.Comparator,
					LHS: res,
					RHS: e,
				}
			}
		}

		return &ExprExpr{Expr: res}, nil
	}

	l, err := toLiteral(rule.Val)
	if err != nil {
		return nil, err
	}

	return &BinaryExpr{
		Op:  rule.Op,
		LHS: &Var{rule.Var},
		RHS: l,
	}, nil
}

func toLiteral(i interface{}) (Expr, error) {
	varType := reflect.TypeOf(i)
	if varType == nil {
		return nil, errors.New("type not supported")
	}

	switch varType.Kind() {
	case reflect.Int:
		return &NumberLiteral{Val: float64(i.(int))}, nil
	case reflect.Int32:
		return &NumberLiteral{Val: float64(i.(int32))}, nil
	case reflect.Int64:
		return &NumberLiteral{Val: float64(i.(int64))}, nil
	case reflect.Float32:
		return &NumberLiteral{Val: float64(i.(float32))}, nil
	case reflect.Float64:
		return &NumberLiteral{Val: float64(i.(float64))}, nil
	case reflect.String:
		return &StringLiteral{Val: i.(string)}, nil
	case reflect.Bool:
		return &BoolLiteral{Val: i.(bool)}, nil
	case reflect.Struct:
		switch i.(type) {
		case time.Time:
			return &TimeLiteral{i.(time.Time)}, nil
		}
	case reflect.Slice:
		switch i.(type) {
		case []string:
			l := &StringSliceLiteral{}
			for _, v := range i.([]string) {
				l.Val = append(l.Val, v)
			}

			return l, nil
		case []int:
			l := &NumberSliceLiteral{}
			for _, v := range i.([]int) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []int32:
			l := &NumberSliceLiteral{}
			for _, v := range i.([]int32) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []int64:
			l := &NumberSliceLiteral{}
			for _, v := range i.([]int64) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []float32:
			l := &NumberSliceLiteral{}
			for _, v := range i.([]float32) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []float64:
			l := &NumberSliceLiteral{}
			for _, v := range i.([]float64) {
				l.Val = append(l.Val, v)
			}

			return l, nil
		case []interface{}:
			arr := i.([]interface{})
			if len(arr) > 0 {
				switch arr[0].(type) {
				case string:
					l := &StringSliceLiteral{}
					for _, v := range arr {
						l.Val = append(l.Val, v.(string))
					}

					return l, nil
				case float64:
					l := &NumberSliceLiteral{}
					for _, v := range arr {
						l.Val = append(l.Val, v.(float64))
					}

					return l, nil
				}
			}
		}
	}

	return nil, errors.New("unsupported literal")
}
