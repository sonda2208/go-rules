package rules

import (
	"encoding/json"
	"errors"
)

type jsonRule struct {
	Comparator Op          `json:"comparator,omitempty"`
	Rules      []jsonRule  `json:"rules,omitempty"`
	Var        string      `json:"var,omitempty"`
	Op         Op          `json:"op,omitempty"`
	Val        interface{} `json:"val,omitempty"`
}

func (r jsonRule) IsValid() error {
	var op Op
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

		op = r.Comparator
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

		op = r.Op
	}

	validOp := false
	ops := []Op{AND, OR, EQ, NEQ, LT, LTE, GT, GTE, IN}
	for _, o := range ops {
		if op == o {
			validOp = true
		}
	}
	if !validOp {
		return errors.New("invalid operator")
	}

	return nil
}

func ParseFromJSON(raw json.RawMessage) (Expr, error) {
	rule := jsonRule{}
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

func parseExpr(rule *jsonRule) (Expr, error) {
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
				res = &binaryExpr{
					Op:  rule.Comparator,
					LHS: res,
					RHS: e,
				}
			}
		}

		return &parentExpr{Expr: res}, nil
	}

	lit, err := toLiteral(rule.Val)
	if err != nil {
		return nil, err
	}

	// parsing "Var" by ANTLR parser to allow users to input custom function
	l, err := ParseFromString(rule.Var)
	if err != nil {
		return nil, err
	}

	return &binaryExpr{
		Op:  rule.Op,
		LHS: l,
		RHS: lit,
	}, nil
}
