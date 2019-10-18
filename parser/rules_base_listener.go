// Code generated from parser/Rules.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rules

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRulesListener is a complete listener for a parse tree produced by RulesParser.
type BaseRulesListener struct{}

var _ RulesListener = &BaseRulesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRulesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRulesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRulesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRulesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseRulesListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseRulesListener) ExitRules(ctx *RulesContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *BaseRulesListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *BaseRulesListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// EnterNumberList is called when production numberList is entered.
func (s *BaseRulesListener) EnterNumberList(ctx *NumberListContext) {}

// ExitNumberList is called when production numberList is exited.
func (s *BaseRulesListener) ExitNumberList(ctx *NumberListContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseRulesListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseRulesListener) ExitStringList(ctx *StringListContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseRulesListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseRulesListener) ExitValueList(ctx *ValueListContext) {}

// EnterIdent is called when production Ident is entered.
func (s *BaseRulesListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production Ident is exited.
func (s *BaseRulesListener) ExitIdent(ctx *IdentContext) {}

// EnterBoolLit is called when production BoolLit is entered.
func (s *BaseRulesListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production BoolLit is exited.
func (s *BaseRulesListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterNumberLit is called when production NumberLit is entered.
func (s *BaseRulesListener) EnterNumberLit(ctx *NumberLitContext) {}

// ExitNumberLit is called when production NumberLit is exited.
func (s *BaseRulesListener) ExitNumberLit(ctx *NumberLitContext) {}

// EnterStringLit is called when production StringLit is entered.
func (s *BaseRulesListener) EnterStringLit(ctx *StringLitContext) {}

// ExitStringLit is called when production StringLit is exited.
func (s *BaseRulesListener) ExitStringLit(ctx *StringLitContext) {}

// EnterDurationLit is called when production DurationLit is entered.
func (s *BaseRulesListener) EnterDurationLit(ctx *DurationLitContext) {}

// ExitDurationLit is called when production DurationLit is exited.
func (s *BaseRulesListener) ExitDurationLit(ctx *DurationLitContext) {}

// EnterFuncExpr is called when production FuncExpr is entered.
func (s *BaseRulesListener) EnterFuncExpr(ctx *FuncExprContext) {}

// ExitFuncExpr is called when production FuncExpr is exited.
func (s *BaseRulesListener) ExitFuncExpr(ctx *FuncExprContext) {}

// EnterInCond is called when production InCond is entered.
func (s *BaseRulesListener) EnterInCond(ctx *InCondContext) {}

// ExitInCond is called when production InCond is exited.
func (s *BaseRulesListener) ExitInCond(ctx *InCondContext) {}

// EnterParentExpr is called when production ParentExpr is entered.
func (s *BaseRulesListener) EnterParentExpr(ctx *ParentExprContext) {}

// ExitParentExpr is called when production ParentExpr is exited.
func (s *BaseRulesListener) ExitParentExpr(ctx *ParentExprContext) {}

// EnterEqualityExpr is called when production EqualityExpr is entered.
func (s *BaseRulesListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production EqualityExpr is exited.
func (s *BaseRulesListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterPrimaryExpr is called when production PrimaryExpr is entered.
func (s *BaseRulesListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production PrimaryExpr is exited.
func (s *BaseRulesListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterRelationalExpr is called when production RelationalExpr is entered.
func (s *BaseRulesListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production RelationalExpr is exited.
func (s *BaseRulesListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterLogicalExpr is called when production LogicalExpr is entered.
func (s *BaseRulesListener) EnterLogicalExpr(ctx *LogicalExprContext) {}

// ExitLogicalExpr is called when production LogicalExpr is exited.
func (s *BaseRulesListener) ExitLogicalExpr(ctx *LogicalExprContext) {}

// EnterBooleanExpr is called when production BooleanExpr is entered.
func (s *BaseRulesListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production BooleanExpr is exited.
func (s *BaseRulesListener) ExitBooleanExpr(ctx *BooleanExprContext) {}
