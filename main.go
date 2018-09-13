package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
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

type Server struct {
	Template *template.Template
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Only GET and POST are supported.", http.StatusMethodNotAllowed)
		return
	}

	var err error
	var expr string
	var result string
	expr = r.FormValue("expr")
	if expr != "" {
		var intResult int
		intResult, err = evaluateExpr(expr)
		result = roman.FromInteger(intResult)
	}

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	data := map[string]string{
		"err":    errMsg,
		"expr":   expr,
		"result": result,
	}

	if err = s.Template.Execute(w, data); err != nil {
		log.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

func main() {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(err)
		return
	}

	server := &Server{
		Template: template,
	}
	http.HandleFunc("/", server.handleIndex)

	log.Info("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
