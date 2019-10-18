package rules

import (
	"fmt"
	"strconv"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sonda2208/go-rules/parser"
)

type exprStack struct {
	stack []Expr
}

func (s *exprStack) Push(v Expr) {
	s.stack = append(s.stack, v)
}

func (s *exprStack) Pop() Expr {
	if len(s.stack) == 0 {
		return nil
	}

	result := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return result
}

func (s *exprStack) Len() int {
	return len(s.stack)
}

type errorListener struct {
	antlr.DefaultErrorListener
	lastError error
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.lastError = fmt.Errorf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

type listener struct {
	*parser.BaseRulesListener
	stack exprStack
}

func (l *listener) ExitNumberList(c *parser.NumberListContext) {
	numbers := &numberSliceLiteral{
		Val: make([]float64, len(c.AllNumberLiteral())),
	}
	for i, n := range c.AllNumberLiteral() {
		val, _ := strconv.ParseFloat(n.GetText(), 64)
		numbers.Val[i] = val
	}

	l.stack.Push(numbers)
}

func (l *listener) ExitStringList(c *parser.StringListContext) {
	strings := &stringSliceLiteral{
		Val: make([]string, len(c.AllStringLiteral())),
	}
	for i, s := range c.AllStringLiteral() {
		strings.Val[i] = unquote(s.GetText())
	}

	l.stack.Push(strings)
}

func (l *listener) ExitIdent(c *parser.IdentContext) {
	l.stack.Push(&identifier{Val: c.GetText()})
}

func (l *listener) ExitBoolLit(c *parser.BoolLitContext) {
	val, _ := strconv.ParseBool(c.GetText())
	l.stack.Push(&boolLiteral{val})
}

func (l *listener) ExitNumberLit(c *parser.NumberLitContext) {
	val, _ := strconv.ParseFloat(c.GetText(), 64)
	l.stack.Push(&numberLiteral{val})
}

func (l *listener) ExitStringLit(c *parser.StringLitContext) {
	val := unquote(c.GetText())

	var expr Expr
	expr = &stringLiteral{val}

	// try parsing to date time
	dt, err := time.Parse(time.RFC3339, val)
	if err == nil {
		expr = &timeLiteral{
			Val: dt,
		}
	}

	l.stack.Push(expr)
}

func (l *listener) EnterDurationLit(c *parser.DurationLitContext) {
	val, _ := time.ParseDuration(c.GetText())
	l.stack.Push(&durationLiteral{val})
}

func (l *listener) ExitFuncExpr(c *parser.FuncExprContext) {
	argListCtx := c.ArgumentExpressionList().(*parser.ArgumentExpressionListContext)
	numOfArgs := len(argListCtx.AllExpression())
	funcExpr := &functionExpression{
		Name:      c.Identifier().GetText(),
		Arguments: make([]Expr, numOfArgs),
	}

	for i := 0; i < numOfArgs; i++ {
		funcExpr.Arguments[i] = l.stack.Pop()
	}

	l.stack.Push(funcExpr)
}

func (l *listener) ExitInCond(c *parser.InCondContext) {
	boolExpr := &binaryExpr{
		Op:  IN,
		LHS: &identifier{c.Identifier().GetText()},
		RHS: l.stack.Pop(),
	}

	l.stack.Push(boolExpr)
}

func (l *listener) ExitParentExpr(c *parser.ParentExprContext) {
	expr := l.stack.Pop()
	l.stack.Push(&parentExpr{
		Expr: expr,
	})
}

func (l *listener) ExitEqualityExpr(c *parser.EqualityExprContext) {
	rhs := l.stack.Pop()
	lhs := l.stack.Pop()

	l.stack.Push(&binaryExpr{
		Op:  Op(c.GetOp().GetText()),
		LHS: lhs,
		RHS: rhs,
	})
}

func (l *listener) ExitRelationalExpr(c *parser.RelationalExprContext) {
	rhs := l.stack.Pop()
	lhs := l.stack.Pop()

	l.stack.Push(&binaryExpr{
		Op:  Op(c.GetOp().GetText()),
		LHS: lhs,
		RHS: rhs,
	})
}

func (l *listener) ExitLogicalExpr(c *parser.LogicalExprContext) {
	rhs := l.stack.Pop()
	lhs := l.stack.Pop()

	l.stack.Push(&binaryExpr{
		Op:  Op(c.GetOp().GetText()),
		LHS: lhs,
		RHS: rhs,
	})
}

func ParseFromString(src string) (Expr, error) {
	e := &errorListener{}

	// make input stream
	inputStream := antlr.NewInputStream(src)

	// create lexer
	lexer := parser.NewRulesLexer(inputStream)
	lexer.RemoveErrorListeners() // remove default lexer error listener
	lexer.AddErrorListener(e)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create parser
	p := parser.NewRulesParser(tokens)
	p.RemoveErrorListeners() // remove default parser error listener
	p.AddErrorListener(e)

	l := &listener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.Rules())

	if e.lastError != nil {
		return nil, e.lastError
	}

	return l.stack.Pop(), nil
}
