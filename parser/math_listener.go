// Code generated from parser/Math.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Math

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MathListener is a complete listener for a parse tree produced by MathParser.
type MathListener interface {
	antlr.ParseTreeListener

	// EnterOrder is called when entering the Order production.
	EnterOrder(c *OrderContext)

	// EnterDiv is called when entering the Div production.
	EnterDiv(c *DivContext)

	// EnterAdd is called when entering the Add production.
	EnterAdd(c *AddContext)

	// EnterSub is called when entering the Sub production.
	EnterSub(c *SubContext)

	// EnterBrace is called when entering the Brace production.
	EnterBrace(c *BraceContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// ExitOrder is called when exiting the Order production.
	ExitOrder(c *OrderContext)

	// ExitDiv is called when exiting the Div production.
	ExitDiv(c *DivContext)

	// ExitAdd is called when exiting the Add production.
	ExitAdd(c *AddContext)

	// ExitSub is called when exiting the Sub production.
	ExitSub(c *SubContext)

	// ExitBrace is called when exiting the Brace production.
	ExitBrace(c *BraceContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)
}
