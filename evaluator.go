package rules

import (
	"errors"
	"fmt"
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

		return toLiteral(varVal)
	}

	return expr, nil
}

func compute(op Op, lhs, rhs Expr) (*BoolLiteral, error) {
	switch op {
	case AND:
		return computeAND(lhs, rhs)
	case OR:
		return computeOR(lhs, rhs)
	case EQ:
		return computeEQ(lhs, rhs)
	case NEQ:
		return computeNEQ(lhs, rhs)
	case LT:
		return computeLT(lhs, rhs)
	case LTE:
		return computeLTE(lhs, rhs)
	case GT:
		return computeGT(lhs, rhs)
	case GTE:
		return computeGTE(lhs, rhs)
	}

	return nil, errors.New("invalid operator")
}

func computeOR(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val || r.Val)}, nil
		}
	}

	return nil, errors.New("unsupported expression")
}

func computeAND(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val && r.Val)}, nil
		}
	}

	return nil, errors.New("unsupported expression")
}

func computeEQ(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val == r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to bool`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}

func computeNEQ(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val != r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val != r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	case *BoolLiteral:
		r, ok := rhs.(*BoolLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val != r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to bool`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}

func computeLT(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val < r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val < r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}

func computeLTE(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val <= r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val <= r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}

func computeGT(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val > r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val > r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}

func computeGTE(lhs, rhs Expr) (*BoolLiteral, error) {
	switch l := lhs.(type) {
	case *NumberLiteral:
		r, ok := rhs.(*NumberLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val >= r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to number`, rhs.String())
	case *StringLiteral:
		r, ok := rhs.(*StringLiteral)
		if ok {
			return &BoolLiteral{Val: (l.Val >= r.Val)}, nil
		}

		return nil, fmt.Errorf(`cannot convert "%s" to string`, rhs.String())
	}

	return nil, errors.New("invalid data type")
}
