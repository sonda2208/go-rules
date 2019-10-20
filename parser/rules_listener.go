// Code generated from parser/Rules.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rules

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RulesListener is a complete listener for a parse tree produced by RulesParser.
type RulesListener interface {
	antlr.ParseTreeListener

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterArgumentExpressionList is called when entering the argumentExpressionList production.
	EnterArgumentExpressionList(c *ArgumentExpressionListContext)

	// EnterNumberList is called when entering the numberList production.
	EnterNumberList(c *NumberListContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterMultiplicativeExpr is called when entering the MultiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterIdent is called when entering the Ident production.
	EnterIdent(c *IdentContext)

	// EnterAdditiveExpr is called when entering the AdditiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterNumberLit is called when entering the NumberLit production.
	EnterNumberLit(c *NumberLitContext)

	// EnterInCond is called when entering the InCond production.
	EnterInCond(c *InCondContext)

	// EnterDurationLit is called when entering the DurationLit production.
	EnterDurationLit(c *DurationLitContext)

	// EnterParentExpr is called when entering the ParentExpr production.
	EnterParentExpr(c *ParentExprContext)

	// EnterFuncExpr is called when entering the FuncExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterBoolLit is called when entering the BoolLit production.
	EnterBoolLit(c *BoolLitContext)

	// EnterStringLit is called when entering the StringLit production.
	EnterStringLit(c *StringLitContext)

	// EnterEqualityExpr is called when entering the EqualityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterRelationalExpr is called when entering the RelationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterLogicalExpr is called when entering the LogicalExpr production.
	EnterLogicalExpr(c *LogicalExprContext)

	// EnterBooleanExpr is called when entering the BooleanExpr production.
	EnterBooleanExpr(c *BooleanExprContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitArgumentExpressionList is called when exiting the argumentExpressionList production.
	ExitArgumentExpressionList(c *ArgumentExpressionListContext)

	// ExitNumberList is called when exiting the numberList production.
	ExitNumberList(c *NumberListContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitMultiplicativeExpr is called when exiting the MultiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitIdent is called when exiting the Ident production.
	ExitIdent(c *IdentContext)

	// ExitAdditiveExpr is called when exiting the AdditiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitNumberLit is called when exiting the NumberLit production.
	ExitNumberLit(c *NumberLitContext)

	// ExitInCond is called when exiting the InCond production.
	ExitInCond(c *InCondContext)

	// ExitDurationLit is called when exiting the DurationLit production.
	ExitDurationLit(c *DurationLitContext)

	// ExitParentExpr is called when exiting the ParentExpr production.
	ExitParentExpr(c *ParentExprContext)

	// ExitFuncExpr is called when exiting the FuncExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitBoolLit is called when exiting the BoolLit production.
	ExitBoolLit(c *BoolLitContext)

	// ExitStringLit is called when exiting the StringLit production.
	ExitStringLit(c *StringLitContext)

	// ExitEqualityExpr is called when exiting the EqualityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitRelationalExpr is called when exiting the RelationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitLogicalExpr is called when exiting the LogicalExpr production.
	ExitLogicalExpr(c *LogicalExprContext)

	// ExitBooleanExpr is called when exiting the BooleanExpr production.
	ExitBooleanExpr(c *BooleanExprContext)
}
