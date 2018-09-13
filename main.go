package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"html/template"
	"math"
	"net/http"
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

func evaluateExpr(expr string) (int, error) {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewMathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMathParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Expr()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	var intResult int
	listener.stack, intResult = listener.stack.Pop()

	return intResult, nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var expr string
	var result string

	expr = r.FormValue("expr")

	if expr != "" {
		intResult, _ := evaluateExpr(expr)
		result = roman.FromInteger(intResult)
	}

	template, err := template.New("index").Parse(`
		<!DOCTYPE html> 
		<html>
			<head>
			</head>
			<body>
				<form>
					Enter expression: <input type="text" name="expr" value="{{ .expr }}">
					<input type="submit" value="submit">
				</form>

				Result: {{ .result }}
			</body>
		</html>`)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Only GET and POST are supported.", http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"expr":   expr,
		"result": result,
	}

	if err = template.Execute(w, data); err != nil {
		fmt.Print(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
