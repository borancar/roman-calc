// Code generated from parser/Math.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Math

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MathListener is a complete listener for a parse tree produced by MathParser.
type MathListener interface {
	antlr.ParseTreeListener

	// EnterOrder is called when entering the Order production.
	EnterOrder(c *OrderContext)

	// EnterBrace is called when entering the Brace production.
	EnterBrace(c *BraceContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// EnterDivMul is called when entering the DivMul production.
	EnterDivMul(c *DivMulContext)

	// ExitOrder is called when exiting the Order production.
	ExitOrder(c *OrderContext)

	// ExitBrace is called when exiting the Brace production.
	ExitBrace(c *BraceContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)

	// ExitDivMul is called when exiting the DivMul production.
	ExitDivMul(c *DivMulContext)
}
