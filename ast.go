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

type Literal interface {
	Value() interface{}
}

type parentExpr struct {
	Expr Expr
}

func (e *parentExpr) String() string {
	return fmt.Sprintf("(%s)", e.Expr.String())
}

func (e *parentExpr) Type() string {
	return "nested expression"
}

type binaryExpr struct {
	Op  Op
	LHS Expr
	RHS Expr
}

func (e *binaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.LHS.String(), e.Op, e.RHS.String())
}

func (e *binaryExpr) Type() string {
	return "binary expression"
}

type identifier struct {
	Val string
}

func (v *identifier) String() string {
	return v.Val
}

func (v *identifier) Type() string {
	return "variable"
}

type numberLiteral struct {
	Val float64
}

func (l *numberLiteral) String() string {
	return strconv.FormatFloat(l.Val, 'f', 3, 64)
}

func (l *numberLiteral) Type() string {
	return "number"
}

func (l *numberLiteral) Value() interface{} {
	return l.Val
}

type stringLiteral struct {
	Val string
}

func (l *stringLiteral) String() string {
	return l.Val
}

func (l *stringLiteral) Type() string {
	return "string"
}

func (l *stringLiteral) Value() interface{} {
	return l.Val
}

type boolLiteral struct {
	Val bool
}

func (l *boolLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *boolLiteral) Type() string {
	return "bool"
}

func (l *boolLiteral) Value() interface{} {
	return l.Val
}

type timeLiteral struct {
	Val time.Time
}

func (l *timeLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *timeLiteral) Type() string {
	return "time"
}

func (l *timeLiteral) Value() interface{} {
	return l.Val
}

type numberSliceLiteral struct {
	Val []float64
}

func (l *numberSliceLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *numberSliceLiteral) Type() string {
	return "number slice"
}

func (l *numberSliceLiteral) Value() interface{} {
	return l.Val
}

type stringSliceLiteral struct {
	Val []string
}

func (l *stringSliceLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func (l *stringSliceLiteral) Type() string {
	return "string slice"
}

func (l *stringSliceLiteral) Value() interface{} {
	return l.Val
}

type durationLiteral struct {
	Val time.Duration
}

func (l *durationLiteral) String() string {
	return l.Val.String()
}

func (l *durationLiteral) Type() string {
	return "duration"
}

func (l *durationLiteral) Value() interface{} {
	return l.Val
}

type Function func(args ...interface{}) (interface{}, error)

type functionExpression struct {
	Name      string
	Arguments []Expr
}

func (l *functionExpression) String() string {
	return l.Name
}

func (l *functionExpression) Type() string {
	return "function"
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
	case *binaryExpr:
		err = walk(e.LHS, fn)
		if err != nil {
			return err
		}

		err = walk(e.RHS, fn)
		if err != nil {
			return err
		}
	case *parentExpr:
		err = walk(e.Expr, fn)
		if err != nil {
			return err
		}
	}

	return nil
}
