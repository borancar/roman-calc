package main

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	loads "github.com/go-openapi/loads"
	middleware "github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"html/template"
	"math"
	"net/http"
	"os"
	"roman-calc/models"
	"roman-calc/parser"
	"roman-calc/restapi"
	"roman-calc/restapi/operations"
	"roman-calc/restapi/operations/calc"
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
	el    *ErrorListener
}

func NewTreeShapeListener(el *ErrorListener) *TreeShapeListener {
	return &TreeShapeListener{
		el: el,
	}
}

func (t *TreeShapeListener) ExitSub(c *parser.SubContext) {
	if len(t.el.Errors) == 0 {
		var op1, op2 int
		t.stack, op2 = t.stack.Pop()
		t.stack, op1 = t.stack.Pop()

		t.stack = t.stack.Push(op1 - op2)
	}
}

func (t *TreeShapeListener) ExitAdd(c *parser.AddContext) {
	if len(t.el.Errors) == 0 {
		var op1, op2 int
		t.stack, op2 = t.stack.Pop()
		t.stack, op1 = t.stack.Pop()

		t.stack = t.stack.Push(op1 + op2)
	}
}

func (t *TreeShapeListener) ExitMul(c *parser.MulContext) {
	if len(t.el.Errors) == 0 {
		var op1, op2 int
		t.stack, op2 = t.stack.Pop()
		t.stack, op1 = t.stack.Pop()

		t.stack = t.stack.Push(op1 * op2)
	}
}

func (t *TreeShapeListener) ExitDiv(c *parser.DivContext) {
	if len(t.el.Errors) == 0 {
		var op1, op2 int
		t.stack, op2 = t.stack.Pop()
		t.stack, op1 = t.stack.Pop()

		t.stack = t.stack.Push(op1 / op2)
	}
}

func (t *TreeShapeListener) ExitOrder(c *parser.OrderContext) {
	if len(t.el.Errors) == 0 {
		var op1, op2 int
		t.stack, op2 = t.stack.Pop()
		t.stack, op1 = t.stack.Pop()

		t.stack = t.stack.Push(int(math.Round(math.Pow(float64(op1), float64(op2)))))
	}
}

func (t *TreeShapeListener) ExitNum(c *parser.NumContext) {
	if len(t.el.Errors) == 0 {
		t.stack = t.stack.Push(roman.ToInteger(c.GetText()))
	}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func NewErrorListener() *ErrorListener {
	return new(ErrorListener)
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Errors = append(el.Errors, errors.New(fmt.Sprintf("Position %d: %s", column, msg)))
}

func evaluateExpr(expr string) (int, []error) {
	errorListener := NewErrorListener()
	treeListener := NewTreeShapeListener(errorListener)

	input := antlr.NewInputStream(expr)
	lexer := parser.NewMathLexer(input)
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMathParser(stream)
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	tree := p.Expr()
	antlr.ParseTreeWalkerDefault.Walk(treeListener, tree)

	var intResult int
	if len(errorListener.Errors) == 0 {
		treeListener.stack, intResult = treeListener.stack.Pop()
	}

	return intResult, errorListener.Errors
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

	var parseErrs []error
	var expr string
	var result string
	expr = r.FormValue("expr")
	if expr != "" {
		var intResult int
		intResult, parseErrs = evaluateExpr(expr)
		result = roman.FromInteger(intResult)
	}

	var errMsgs []string
	for _, e := range parseErrs {
		errMsgs = append(errMsgs, e.Error())
	}

	data := map[string]interface{}{
		"errs":   errMsgs,
		"expr":   expr,
		"result": result,
	}

	if err := s.Template.Execute(w, data); err != nil {
		log.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

func main() {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	server := &Server{
		Template: template,
	}
	http.HandleFunc("/", server.handleIndex)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	api := operations.NewRomanCalculatorAPI(swaggerSpec)
	api.CalcCalcExprHandler = calc.CalcExprHandlerFunc(func(params calc.CalcExprParams) middleware.Responder {
		intResult, errs := evaluateExpr(params.Expr)
		data := &models.Result{}
		for _, err := range errs {
			data.Errs = append(data.Errs, err.Error())
		}
		data.Expr = params.Expr
		data.Result = roman.FromInteger(intResult)
		return calc.NewCalcExprOK().WithPayload(data)
	})
	http.Handle("/calc", api.Serve(nil))

	log.Info("Listening on :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
