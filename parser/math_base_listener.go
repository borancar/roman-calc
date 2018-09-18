// Code generated from parser/Math.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Math

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMathListener is a complete listener for a parse tree produced by MathParser.
type BaseMathListener struct{}

var _ MathListener = &BaseMathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOrder is called when production Order is entered.
func (s *BaseMathListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production Order is exited.
func (s *BaseMathListener) ExitOrder(ctx *OrderContext) {}

// EnterBrace is called when production Brace is entered.
func (s *BaseMathListener) EnterBrace(ctx *BraceContext) {}

// ExitBrace is called when production Brace is exited.
func (s *BaseMathListener) ExitBrace(ctx *BraceContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseMathListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseMathListener) ExitAddSub(ctx *AddSubContext) {}

// EnterNum is called when production Num is entered.
func (s *BaseMathListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production Num is exited.
func (s *BaseMathListener) ExitNum(ctx *NumContext) {}

// EnterDivMul is called when production DivMul is entered.
func (s *BaseMathListener) EnterDivMul(ctx *DivMulContext) {}

// ExitDivMul is called when production DivMul is exited.
func (s *BaseMathListener) ExitDivMul(ctx *DivMulContext) {}
