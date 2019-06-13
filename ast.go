package rules

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Expr interface {
	String() string
	Type() string
}

type ParentExpr struct {
	Expr Expr
}

func (e *ParentExpr) String() string {
	return fmt.Sprintf("(%s)", e.Expr.String())
}

func (e *ParentExpr) Type() string {
	return "nested expression"
}

type BinaryExpr struct {
	Op  Op
	LHS Expr
	RHS Expr
}

func (e *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.LHS.String(), e.Op, e.RHS.String())
}

func (e *BinaryExpr) Type() string {
	return "binary expression"
}

type Ident struct {
	Val string
}

func (v *Ident) String() string {
	return v.Val
}

func (v *Ident) Type() string {
	return "variable"
}

type NumberLiteral struct {
	Val float64
}

func (l *NumberLiteral) String() string {
	return strconv.FormatFloat(l.Val, 'f', 3, 64)
}

func (l *NumberLiteral) Type() string {
	return "number"
}

type StringLiteral struct {
	Val string
}

func (l *StringLiteral) String() string {
	return l.Val
}

func (l *StringLiteral) Type() string {
	return "string"
}

type BoolLiteral struct {
	Val bool
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *BoolLiteral) Type() string {
	return "bool"
}

type TimeLiteral struct {
	Val time.Time
}

func (l *TimeLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *TimeLiteral) Type() string {
	return "time"
}

type NumberSliceLiteral struct {
	Val []float64
}

func (l *NumberSliceLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *NumberSliceLiteral) Type() string {
	return "number slice"
}

type StringSliceLiteral struct {
	Val []string
}

func (l *StringSliceLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *StringSliceLiteral) Type() string {
	return "string slice"
}

type DurationLiteral struct {
	Val time.Duration
}

func (l *DurationLiteral) String() string {
	return l.Val.String()
}

func (l *DurationLiteral) Type() string {
	return "duration"
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
	case *ParentExpr:
		err = walk(e.Expr, fn)
		if err != nil {
			return err
		}
	}

	return nil
}
