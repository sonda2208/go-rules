package rules

import (
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
	return e.String()
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
