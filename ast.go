package rules

import (
	"errors"
	"fmt"
	"strconv"
)

type Expr interface {
	String() string
}

type ExprExpr struct {
	Expr Expr
}

func (e *ExprExpr) String() string {
	return fmt.Sprintf("(%s)", e.Expr.String())
}

type BinaryExpr struct {
	Op  Op
	LHS Expr
	RHS Expr
}

func (e *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.LHS.String(), e.Op, e.RHS.String())
}

type Var struct {
	Val string
}

func (v *Var) String() string {
	return v.Val
}

type NumberLiteral struct {
	Val float64
}

func (l *NumberLiteral) String() string {
	return strconv.FormatFloat(l.Val, 'f', 3, 64)
}

type StringLiteral struct {
	Val string
}

func (l *StringLiteral) String() string {
	return l.Val
}

type BoolLiteral struct {
	Val bool
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

type WalkFunc func(expr Expr, err error) error

func Walk(expr Expr, fn WalkFunc) error {
	if expr == nil {
		return fn(nil, errors.New("expression must not be nil"))
	}

	return walk(expr, fn)
}

func walk(expr Expr, fn WalkFunc) error {
	if expr == nil {
		return nil
	}

	err := fn(expr, nil)
	if err != nil {
		return err
	}

	switch e := expr.(type) {
	case *BinaryExpr:
		err = walk(e.LHS, fn)
		if err != nil {
			return err
		}

		err = walk(e.RHS, fn)
		if err != nil {
			return err
		}

	case *ExprExpr:
		err = walk(e.Expr, fn)
		if err != nil {
			return err
		}
	}

	return nil
}
