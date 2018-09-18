// Code generated from parser/Math.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Math

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 27, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14,
	2, 25, 11, 2, 3, 2, 2, 3, 2, 3, 2, 2, 4, 3, 2, 5, 6, 3, 2, 3, 4, 2, 29,
	2, 10, 3, 2, 2, 2, 4, 5, 8, 2, 1, 2, 5, 6, 7, 9, 2, 2, 6, 7, 5, 2, 2, 2,
	7, 8, 7, 10, 2, 2, 8, 11, 3, 2, 2, 2, 9, 11, 7, 8, 2, 2, 10, 4, 3, 2, 2,
	2, 10, 9, 3, 2, 2, 2, 11, 23, 3, 2, 2, 2, 12, 13, 12, 6, 2, 2, 13, 14,
	7, 7, 2, 2, 14, 22, 5, 2, 2, 7, 15, 16, 12, 5, 2, 2, 16, 17, 9, 2, 2, 2,
	17, 22, 5, 2, 2, 6, 18, 19, 12, 4, 2, 2, 19, 20, 9, 3, 2, 2, 20, 22, 5,
	2, 2, 5, 21, 12, 3, 2, 2, 2, 21, 15, 3, 2, 2, 2, 21, 18, 3, 2, 2, 2, 22,
	25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2,
	2, 25, 23, 3, 2, 2, 2, 5, 10, 21, 23,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'^'", "", "'('", "')'",
}
var symbolicNames = []string{
	"", "ADD", "SUB", "MUL", "DIV", "POW", "ROMAN", "LPAREN", "RPAREN", "WS",
}

var ruleNames = []string{
	"expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MathParser struct {
	*antlr.BaseParser
}

func NewMathParser(input antlr.TokenStream) *MathParser {
	this := new(MathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Math.g4"

	return this
}

// MathParser tokens.
const (
	MathParserEOF    = antlr.TokenEOF
	MathParserADD    = 1
	MathParserSUB    = 2
	MathParserMUL    = 3
	MathParserDIV    = 4
	MathParserPOW    = 5
	MathParserROMAN  = 6
	MathParserLPAREN = 7
	MathParserRPAREN = 8
	MathParserWS     = 9
)

// MathParserRULE_expr is the MathParser rule.
const MathParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MathParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrderContext struct {
	*ExprContext
}

func NewOrderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrderContext {
	var p = new(OrderContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrderContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrderContext) POW() antlr.TerminalNode {
	return s.GetToken(MathParserPOW, 0)
}

func (s *OrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOrder(s)
	}
}

func (s *OrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOrder(s)
	}
}

type BraceContext struct {
	*ExprContext
}

func NewBraceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BraceContext {
	var p = new(BraceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MathParserLPAREN, 0)
}

func (s *BraceContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BraceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MathParserRPAREN, 0)
}

func (s *BraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterBrace(s)
	}
}

func (s *BraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitBrace(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(MathParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(MathParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type NumContext struct {
	*ExprContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ROMAN() antlr.TerminalNode {
	return s.GetToken(MathParserROMAN, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitNum(s)
	}
}

type DivMulContext struct {
	*ExprContext
	op antlr.Token
}

func NewDivMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivMulContext {
	var p = new(DivMulContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivMulContext) GetOp() antlr.Token { return s.op }

func (s *DivMulContext) SetOp(v antlr.Token) { s.op = v }

func (s *DivMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivMulContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivMulContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivMulContext) DIV() antlr.TerminalNode {
	return s.GetToken(MathParserDIV, 0)
}

func (s *DivMulContext) MUL() antlr.TerminalNode {
	return s.GetToken(MathParserMUL, 0)
}

func (s *DivMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterDivMul(s)
	}
}

func (s *DivMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitDivMul(s)
	}
}

func (p *MathParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MathParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, MathParserRULE_expr, _p)
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
	p.SetState(8)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MathParserLPAREN:
		localctx = NewBraceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(MathParserLPAREN)
		}
		{
			p.SetState(4)
			p.expr(0)
		}
		{
			p.SetState(5)
			p.Match(MathParserRPAREN)
		}

	case MathParserROMAN:
		localctx = NewNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(7)
			p.Match(MathParserROMAN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOrderContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MathParserRULE_expr)
				p.SetState(10)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(11)
					p.Match(MathParserPOW)
				}
				{
					p.SetState(12)
					p.expr(5)
				}

			case 2:
				localctx = NewDivMulContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MathParserRULE_expr)
				p.SetState(13)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(14)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*DivMulContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MathParserMUL || _la == MathParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*DivMulContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(15)
					p.expr(4)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MathParserRULE_expr)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MathParserADD || _la == MathParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)
					p.expr(3)
				}

			}

		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *MathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MathParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
