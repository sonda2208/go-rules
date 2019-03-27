package rules

import (
	"encoding/json"
	"errors"
	"reflect"
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

	typeof := reflect.TypeOf(rule.Val)
	if typeof == nil {
		return nil, errors.New("invalid type")
	}

	switch typeof.Kind() {
	case reflect.Int:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &NumberLiteral{Val: float64(rule.Val.(int))},
		}, nil
	case reflect.Int32:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &NumberLiteral{Val: float64(rule.Val.(int32))},
		}, nil
	case reflect.Int64:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &NumberLiteral{Val: float64(rule.Val.(int64))},
		}, nil
	case reflect.Float32:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &NumberLiteral{Val: float64(rule.Val.(float32))},
		}, nil
	case reflect.Float64:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &NumberLiteral{Val: float64(rule.Val.(float64))},
		}, nil
	case reflect.String:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &StringLiteral{Val: rule.Val.(string)},
		}, nil
	case reflect.Bool:
		return &BinaryExpr{
			Op:  rule.Op,
			LHS: &Var{rule.Var},
			RHS: &BoolLiteral{Val: rule.Val.(bool)},
		}, nil
	}

	return nil, errors.New("unexpected error")
}
