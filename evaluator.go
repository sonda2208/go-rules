package rules

import (
	"errors"
	"fmt"
	"reflect"
)

func Evaluate(expr Expr, params map[string]interface{}) (bool, error) {
	if expr == nil {
		return false, errors.New("invalid input expression")
	}

	res, err := evaluate(expr, params)
	if err != nil {
		return false, err
	}

	l, ok := res.(*BoolLiteral)
	if !ok {
		return false, errors.New("invalid result expression")
	}

	return l.Val, nil
}

func evaluate(expr Expr, params map[string]interface{}) (Expr, error) {
	if expr == nil {
		return nil, errors.New("invalid expression")
	}

	switch e := expr.(type) {
	case *ExprExpr:
		return evaluate(e.Expr, params)
	case *BinaryExpr:
		l, err := evaluate(e.LHS, params)
		if err != nil {
			return nil, err
		}

		r, err := evaluate(e.RHS, params)
		if err != nil {
			return nil, err
		}

		return compute(e.Op, l, r)
	case *Var:
		varName := e.Val
		varVal, ok := params[varName]
		if !ok {
			return nil, fmt.Errorf("param %s not found", varName)
		}

		varType := reflect.TypeOf(params[varName])
		if varType == nil {
			return nil, errors.New("type not supported")
		}

		switch varType.Kind() {
		case reflect.Int:
			return &NumberLiteral{Val: float64(varVal.(int))}, nil
		case reflect.Int32:
			return &NumberLiteral{Val: float64(varVal.(int32))}, nil
		case reflect.Int64:
			return &NumberLiteral{Val: float64(varVal.(int64))}, nil
		case reflect.Float32:
			return &NumberLiteral{Val: float64(varVal.(float32))}, nil
		case reflect.Float64:
			return &NumberLiteral{Val: float64(varVal.(float64))}, nil
		case reflect.String:
			return &StringLiteral{Val: varVal.(string)}, nil
		case reflect.Bool:
			return &BoolLiteral{Val: varVal.(bool)}, nil
		}
	}

	return expr, nil
}

func compute(op Op, lhs, rhs Expr) (*BoolLiteral, error) {
	switch op {
	case AND:
		return computeAnd(lhs, rhs)
	case OR:
		return computeOr(lhs, rhs)
	case EQ:
		return computeEq(lhs, rhs)
	}

	return nil, nil
}

func computeOr(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val || r.Val)}, nil
		}
	}

	return nil, errors.New("unsupported expression")
}

func computeAnd(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val && r.Val)}, nil
		}
	}

	return nil, errors.New("unsupported expression")
}

func computeEq(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}
	}

	return nil, errors.New("unsupported expression")
}
