package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"os"
	"roman-calc/parser"
	"roman-calc/roman"
)

type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, int) {
	l := len(s)

	return s[:l-1], s[l-1]
}

type TreeShapeListener struct {
	*parser.BaseMathListener
	stack Stack
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) ExitSub(c *parser.SubContext) {
	var op1, op2 int
	t.stack, op2 = t.stack.Pop()
	t.stack, op1 = t.stack.Pop()

	t.stack = t.stack.Push(op1 - op2)
}

func (t *TreeShapeListener) ExitAdd(c *parser.AddContext) {
	var op1, op2 int
	t.stack, op2 = t.stack.Pop()
	t.stack, op1 = t.stack.Pop()

	t.stack = t.stack.Push(op1 + op2)
}

func (t *TreeShapeListener) ExitMul(c *parser.MulContext) {
	var op1, op2 int
	t.stack, op2 = t.stack.Pop()
	t.stack, op1 = t.stack.Pop()

	t.stack = t.stack.Push(op1 * op2)
}

func (t *TreeShapeListener) ExitDiv(c *parser.DivContext) {
	var op1, op2 int
	t.stack, op2 = t.stack.Pop()
	t.stack, op1 = t.stack.Pop()

	t.stack = t.stack.Push(op1 / op2)
}

func (t *TreeShapeListener) ExitOrder(c *parser.OrderContext) {
	var op1, op2 int
	t.stack, op2 = t.stack.Pop()
	t.stack, op1 = t.stack.Pop()

	t.stack = t.stack.Push(int(math.Round(math.Pow(float64(op1), float64(op2)))))
}

func (t *TreeShapeListener) ExitNum(c *parser.NumContext) {
	t.stack = append(t.stack, roman.ToInteger(c.GetText()))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter expression: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		input := antlr.NewInputStream(text)
		lexer := parser.NewMathLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewMathParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Expr()
		listener := NewTreeShapeListener()
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)

		var res int
		listener.stack, res = listener.stack.Pop()
		fmt.Printf("Result: %s\n", roman.FromInteger(res))
	}
}
