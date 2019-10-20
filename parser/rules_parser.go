// Code generated from parser/Rules.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 119,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 5, 3, 31, 10, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3, 5, 3, 5, 3, 5,
	7, 5, 46, 10, 5, 12, 5, 14, 5, 49, 11, 5, 3, 6, 3, 6, 6, 6, 53, 10, 6,
	13, 6, 14, 6, 54, 3, 6, 6, 6, 58, 10, 6, 13, 6, 14, 6, 59, 5, 6, 62, 10,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 81, 10, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 89, 10, 7, 12, 7, 14, 7, 92, 11, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 103, 10, 8, 12, 8, 14, 8,
	106, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 114, 10, 9, 12, 9,
	14, 9, 117, 11, 9, 3, 9, 2, 5, 12, 14, 16, 10, 2, 4, 6, 8, 10, 12, 14,
	16, 2, 7, 3, 2, 17, 19, 3, 2, 15, 16, 3, 2, 8, 9, 3, 2, 10, 13, 3, 2, 6,
	7, 2, 129, 2, 18, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8,
	42, 3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 80, 3, 2, 2, 2, 14, 93, 3, 2, 2,
	2, 16, 107, 3, 2, 2, 2, 18, 19, 5, 16, 9, 2, 19, 20, 7, 2, 2, 3, 20, 3,
	3, 2, 2, 2, 21, 30, 7, 20, 2, 2, 22, 27, 5, 16, 9, 2, 23, 24, 7, 3, 2,
	2, 24, 26, 5, 16, 9, 2, 25, 23, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25,
	3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2,
	30, 22, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 7,
	21, 2, 2, 33, 5, 3, 2, 2, 2, 34, 39, 7, 23, 2, 2, 35, 36, 7, 3, 2, 2, 36,
	38, 7, 23, 2, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2,
	2, 2, 39, 40, 3, 2, 2, 2, 40, 7, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 47,
	7, 25, 2, 2, 43, 44, 7, 3, 2, 2, 44, 46, 7, 25, 2, 2, 45, 43, 3, 2, 2,
	2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 9, 3,
	2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 61, 7, 4, 2, 2, 51, 53, 5, 6, 4, 2, 52,
	51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2,
	2, 55, 62, 3, 2, 2, 2, 56, 58, 5, 8, 5, 2, 57, 56, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2,
	61, 52, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7,
	5, 2, 2, 64, 11, 3, 2, 2, 2, 65, 66, 8, 7, 1, 2, 66, 81, 7, 26, 2, 2, 67,
	81, 7, 22, 2, 2, 68, 81, 7, 23, 2, 2, 69, 81, 7, 25, 2, 2, 70, 81, 7, 24,
	2, 2, 71, 72, 7, 26, 2, 2, 72, 73, 7, 14, 2, 2, 73, 81, 5, 10, 6, 2, 74,
	75, 7, 26, 2, 2, 75, 81, 5, 4, 3, 2, 76, 77, 7, 20, 2, 2, 77, 78, 5, 16,
	9, 2, 78, 79, 7, 21, 2, 2, 79, 81, 3, 2, 2, 2, 80, 65, 3, 2, 2, 2, 80,
	67, 3, 2, 2, 2, 80, 68, 3, 2, 2, 2, 80, 69, 3, 2, 2, 2, 80, 70, 3, 2, 2,
	2, 80, 71, 3, 2, 2, 2, 80, 74, 3, 2, 2, 2, 80, 76, 3, 2, 2, 2, 81, 90,
	3, 2, 2, 2, 82, 83, 12, 4, 2, 2, 83, 84, 9, 2, 2, 2, 84, 89, 5, 12, 7,
	5, 85, 86, 12, 3, 2, 2, 86, 87, 9, 3, 2, 2, 87, 89, 5, 12, 7, 4, 88, 82,
	3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2,
	90, 91, 3, 2, 2, 2, 91, 13, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 8,
	8, 1, 2, 94, 95, 5, 12, 7, 2, 95, 104, 3, 2, 2, 2, 96, 97, 12, 4, 2, 2,
	97, 98, 9, 4, 2, 2, 98, 103, 5, 12, 7, 2, 99, 100, 12, 3, 2, 2, 100, 101,
	9, 5, 2, 2, 101, 103, 5, 12, 7, 2, 102, 96, 3, 2, 2, 2, 102, 99, 3, 2,
	2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2,
	105, 15, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 108, 8, 9, 1, 2, 108, 109,
	5, 14, 8, 2, 109, 115, 3, 2, 2, 2, 110, 111, 12, 3, 2, 2, 111, 112, 9,
	6, 2, 2, 112, 114, 5, 14, 8, 2, 113, 110, 3, 2, 2, 2, 114, 117, 3, 2, 2,
	2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 17, 3, 2, 2, 2, 117,
	115, 3, 2, 2, 2, 15, 27, 30, 39, 47, 54, 59, 61, 80, 88, 90, 102, 104,
	115,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'['", "']'", "'&&'", "'||'", "'=='", "'!='", "'<'", "'<='",
	"'>'", "'>='", "'in'", "'+'", "'-'", "'*'", "'/'", "'%'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT", "GTE", "IN",
	"ADD", "SUB", "MUL", "DIV", "MOD", "LPARENT", "RPARENT", "BooleanLiteral",
	"NumberLiteral", "DurationLiteral", "StringLiteral", "Identifier", "WS",
}

var ruleNames = []string{
	"rules", "argumentExpressionList", "numberList", "stringList", "valueList",
	"primaryExpression", "booleanExpression", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RulesParser struct {
	*antlr.BaseParser
}

func NewRulesParser(input antlr.TokenStream) *RulesParser {
	this := new(RulesParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rules.g4"

	return this
}

// RulesParser tokens.
const (
	RulesParserEOF             = antlr.TokenEOF
	RulesParserT__0            = 1
	RulesParserT__1            = 2
	RulesParserT__2            = 3
	RulesParserAND             = 4
	RulesParserOR              = 5
	RulesParserEQ              = 6
	RulesParserNEQ             = 7
	RulesParserLT              = 8
	RulesParserLTE             = 9
	RulesParserGT              = 10
	RulesParserGTE             = 11
	RulesParserIN              = 12
	RulesParserADD             = 13
	RulesParserSUB             = 14
	RulesParserMUL             = 15
	RulesParserDIV             = 16
	RulesParserMOD             = 17
	RulesParserLPARENT         = 18
	RulesParserRPARENT         = 19
	RulesParserBooleanLiteral  = 20
	RulesParserNumberLiteral   = 21
	RulesParserDurationLiteral = 22
	RulesParserStringLiteral   = 23
	RulesParserIdentifier      = 24
	RulesParserWS              = 25
)

// RulesParser rules.
const (
	RulesParserRULE_rules                  = 0
	RulesParserRULE_argumentExpressionList = 1
	RulesParserRULE_numberList             = 2
	RulesParserRULE_stringList             = 3
	RulesParserRULE_valueList              = 4
	RulesParserRULE_primaryExpression      = 5
	RulesParserRULE_booleanExpression      = 6
	RulesParserRULE_expression             = 7
)

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RulesContext) EOF() antlr.TerminalNode {
	return s.GetToken(RulesParserEOF, 0)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitRules(s)
	}
}

func (p *RulesParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RulesParserRULE_rules)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expression(0)
	}
	{
		p.SetState(17)
		p.Match(RulesParserEOF)
	}

	return localctx
}

// IArgumentExpressionListContext is an interface to support dynamic dispatch.
type IArgumentExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentExpressionListContext differentiates from other interfaces.
	IsArgumentExpressionListContext()
}

type ArgumentExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentExpressionListContext() *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_argumentExpressionList
	return p
}

func (*ArgumentExpressionListContext) IsArgumentExpressionListContext() {}

func NewArgumentExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_argumentExpressionList

	return p
}

func (s *ArgumentExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentExpressionListContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserLPARENT, 0)
}

func (s *ArgumentExpressionListContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserRPARENT, 0)
}

func (s *ArgumentExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterArgumentExpressionList(s)
	}
}

func (s *ArgumentExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitArgumentExpressionList(s)
	}
}

func (p *RulesParser) ArgumentExpressionList() (localctx IArgumentExpressionListContext) {
	localctx = NewArgumentExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RulesParserRULE_argumentExpressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(RulesParserLPARENT)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RulesParserLPARENT)|(1<<RulesParserBooleanLiteral)|(1<<RulesParserNumberLiteral)|(1<<RulesParserDurationLiteral)|(1<<RulesParserStringLiteral)|(1<<RulesParserIdentifier))) != 0 {
		{
			p.SetState(20)
			p.expression(0)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RulesParserT__0 {
			{
				p.SetState(21)
				p.Match(RulesParserT__0)
			}
			{
				p.SetState(22)
				p.expression(0)
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(30)
		p.Match(RulesParserRPARENT)
	}

	return localctx
}

// INumberListContext is an interface to support dynamic dispatch.
type INumberListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberListContext differentiates from other interfaces.
	IsNumberListContext()
}

type NumberListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberListContext() *NumberListContext {
	var p = new(NumberListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_numberList
	return p
}

func (*NumberListContext) IsNumberListContext() {}

func NewNumberListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberListContext {
	var p = new(NumberListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_numberList

	return p
}

func (s *NumberListContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberListContext) AllNumberLiteral() []antlr.TerminalNode {
	return s.GetTokens(RulesParserNumberLiteral)
}

func (s *NumberListContext) NumberLiteral(i int) antlr.TerminalNode {
	return s.GetToken(RulesParserNumberLiteral, i)
}

func (s *NumberListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterNumberList(s)
	}
}

func (s *NumberListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitNumberList(s)
	}
}

func (p *RulesParser) NumberList() (localctx INumberListContext) {
	localctx = NewNumberListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RulesParserRULE_numberList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(RulesParserNumberLiteral)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RulesParserT__0 {
		{
			p.SetState(33)
			p.Match(RulesParserT__0)
		}
		{
			p.SetState(34)
			p.Match(RulesParserNumberLiteral)
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStringListContext is an interface to support dynamic dispatch.
type IStringListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringListContext differentiates from other interfaces.
	IsStringListContext()
}

type StringListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListContext() *StringListContext {
	var p = new(StringListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_stringList
	return p
}

func (*StringListContext) IsStringListContext() {}

func NewStringListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListContext {
	var p = new(StringListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_stringList

	return p
}

func (s *StringListContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(RulesParserStringLiteral)
}

func (s *StringListContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(RulesParserStringLiteral, i)
}

func (s *StringListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterStringList(s)
	}
}

func (s *StringListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitStringList(s)
	}
}

func (p *RulesParser) StringList() (localctx IStringListContext) {
	localctx = NewStringListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RulesParserRULE_stringList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(RulesParserStringLiteral)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RulesParserT__0 {
		{
			p.SetState(41)
			p.Match(RulesParserT__0)
		}
		{
			p.SetState(42)
			p.Match(RulesParserStringLiteral)
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_valueList
	return p
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) AllNumberList() []INumberListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberListContext)(nil)).Elem())
	var tst = make([]INumberListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberListContext)
		}
	}

	return tst
}

func (s *ValueListContext) NumberList(i int) INumberListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberListContext)
}

func (s *ValueListContext) AllStringList() []IStringListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringListContext)(nil)).Elem())
	var tst = make([]IStringListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringListContext)
		}
	}

	return tst
}

func (s *ValueListContext) StringList(i int) IStringListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringListContext)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (p *RulesParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RulesParserRULE_valueList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(RulesParserT__1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RulesParserNumberLiteral:
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == RulesParserNumberLiteral {
			{
				p.SetState(49)
				p.NumberList()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case RulesParserStringLiteral:
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == RulesParserStringLiteral {
			{
				p.SetState(54)
				p.StringList()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(61)
		p.Match(RulesParserT__2)
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyFrom(ctx *PrimaryExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiplicativeExprContext struct {
	*PrimaryExpressionContext
	op antlr.Token
}

func NewMultiplicativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *MultiplicativeExprContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExprContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *MultiplicativeExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(RulesParserMUL, 0)
}

func (s *MultiplicativeExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(RulesParserDIV, 0)
}

func (s *MultiplicativeExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(RulesParserMOD, 0)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

type IdentContext struct {
	*PrimaryExpressionContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitIdent(s)
	}
}

type AdditiveExprContext struct {
	*PrimaryExpressionContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExprContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *AdditiveExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(RulesParserADD, 0)
}

func (s *AdditiveExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(RulesParserSUB, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

type NumberLitContext struct {
	*PrimaryExpressionContext
}

func NewNumberLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLitContext {
	var p = new(NumberLitContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *NumberLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLitContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserNumberLiteral, 0)
}

func (s *NumberLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterNumberLit(s)
	}
}

func (s *NumberLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitNumberLit(s)
	}
}

type InCondContext struct {
	*PrimaryExpressionContext
}

func NewInCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InCondContext {
	var p = new(InCondContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *InCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InCondContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *InCondContext) IN() antlr.TerminalNode {
	return s.GetToken(RulesParserIN, 0)
}

func (s *InCondContext) ValueList() IValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *InCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterInCond(s)
	}
}

func (s *InCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitInCond(s)
	}
}

type DurationLitContext struct {
	*PrimaryExpressionContext
}

func NewDurationLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DurationLitContext {
	var p = new(DurationLitContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *DurationLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationLitContext) DurationLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserDurationLiteral, 0)
}

func (s *DurationLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterDurationLit(s)
	}
}

func (s *DurationLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitDurationLit(s)
	}
}

type ParentExprContext struct {
	*PrimaryExpressionContext
}

func NewParentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentExprContext {
	var p = new(ParentExprContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *ParentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentExprContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserLPARENT, 0)
}

func (s *ParentExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParentExprContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserRPARENT, 0)
}

func (s *ParentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterParentExpr(s)
	}
}

func (s *ParentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitParentExpr(s)
	}
}

type FuncExprContext struct {
	*PrimaryExpressionContext
}

func NewFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExprContext {
	var p = new(FuncExprContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *FuncExprContext) ArgumentExpressionList() IArgumentExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentExpressionListContext)
}

func (s *FuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterFuncExpr(s)
	}
}

func (s *FuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitFuncExpr(s)
	}
}

type BoolLitContext struct {
	*PrimaryExpressionContext
}

func NewBoolLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLitContext {
	var p = new(BoolLitContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserBooleanLiteral, 0)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

type StringLitContext struct {
	*PrimaryExpressionContext
}

func NewStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLitContext {
	var p = new(StringLitContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *StringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLitContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserStringLiteral, 0)
}

func (s *StringLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterStringLit(s)
	}
}

func (s *StringLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitStringLit(s)
	}
}

func (p *RulesParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *RulesParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, RulesParserRULE_primaryExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(64)
			p.Match(RulesParserIdentifier)
		}

	case 2:
		localctx = NewBoolLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(RulesParserBooleanLiteral)
		}

	case 3:
		localctx = NewNumberLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(RulesParserNumberLiteral)
		}

	case 4:
		localctx = NewStringLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(RulesParserStringLiteral)
		}

	case 5:
		localctx = NewDurationLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(RulesParserDurationLiteral)
		}

	case 6:
		localctx = NewInCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(RulesParserIdentifier)
		}
		{
			p.SetState(70)
			p.Match(RulesParserIN)
		}
		{
			p.SetState(71)
			p.ValueList()
		}

	case 7:
		localctx = NewFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(RulesParserIdentifier)
		}
		{
			p.SetState(73)
			p.ArgumentExpressionList()
		}

	case 8:
		localctx = NewParentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(RulesParserLPARENT)
		}
		{
			p.SetState(75)
			p.expression(0)
		}
		{
			p.SetState(76)
			p.Match(RulesParserRPARENT)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExprContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_primaryExpression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(81)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicativeExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RulesParserMUL)|(1<<RulesParserDIV)|(1<<RulesParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicativeExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)
					p.primaryExpression(3)
				}

			case 2:
				localctx = NewAdditiveExprContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_primaryExpression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(84)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RulesParserADD || _la == RulesParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(85)
					p.primaryExpression(2)
				}

			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanExpressionContext is an interface to support dynamic dispatch.
type IBooleanExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanExpressionContext differentiates from other interfaces.
	IsBooleanExpressionContext()
}

type BooleanExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanExpressionContext() *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_booleanExpression
	return p
}

func (*BooleanExpressionContext) IsBooleanExpressionContext() {}

func NewBooleanExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_booleanExpression

	return p
}

func (s *BooleanExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExpressionContext) CopyFrom(ctx *BooleanExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualityExprContext struct {
	*BooleanExpressionContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *EqualityExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(RulesParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(RulesParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

type PrimaryExprContext struct {
	*BooleanExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

type RelationalExprContext struct {
	*BooleanExpressionContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *RelationalExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(RulesParserLT, 0)
}

func (s *RelationalExprContext) LTE() antlr.TerminalNode {
	return s.GetToken(RulesParserLTE, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(RulesParserGT, 0)
}

func (s *RelationalExprContext) GTE() antlr.TerminalNode {
	return s.GetToken(RulesParserGTE, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (p *RulesParser) BooleanExpression() (localctx IBooleanExpressionContext) {
	return p.booleanExpression(0)
}

func (p *RulesParser) booleanExpression(_p int) (localctx IBooleanExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, RulesParserRULE_booleanExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPrimaryExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(92)
		p.primaryExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualityExprContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_booleanExpression)
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(95)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RulesParserEQ || _la == RulesParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(96)
					p.primaryExpression(0)
				}

			case 2:
				localctx = NewRelationalExprContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_booleanExpression)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(98)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RulesParserLT)|(1<<RulesParserLTE)|(1<<RulesParserGT)|(1<<RulesParserGTE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(99)
					p.primaryExpression(0)
				}

			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RulesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LogicalExprContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewLogicalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExprContext {
	var p = new(LogicalExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExprContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *LogicalExprContext) AND() antlr.TerminalNode {
	return s.GetToken(RulesParserAND, 0)
}

func (s *LogicalExprContext) OR() antlr.TerminalNode {
	return s.GetToken(RulesParserOR, 0)
}

func (s *LogicalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterLogicalExpr(s)
	}
}

func (s *LogicalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitLogicalExpr(s)
	}
}

type BooleanExprContext struct {
	*ExpressionContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *BooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterBooleanExpr(s)
	}
}

func (s *BooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitBooleanExpr(s)
	}
}

func (p *RulesParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RulesParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, RulesParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewBooleanExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(106)
		p.booleanExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_expression)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(109)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*LogicalExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == RulesParserAND || _la == RulesParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*LogicalExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(110)
				p.booleanExpression(0)
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

func (p *RulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	case 6:
		var t *BooleanExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanExpressionContext)
		}
		return p.BooleanExpression_Sempred(t, predIndex)

	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RulesParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RulesParser) BooleanExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RulesParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
