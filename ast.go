package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
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

type TimeLiteral struct {
	Val time.Time
}

func (l *TimeLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

type NumberSliceLiteral struct {
	Val []float64
}

func (l *NumberSliceLiteral) String() string {
	return fmt.Sprintf("%v", l.Val)
}

type StringSliceLiteral struct {
	Val []string
}

func (l *StringSliceLiteral) String() string {
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
		}
	}

	return nil, nil
}
