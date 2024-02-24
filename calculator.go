package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
)

// CalculatorModule handles basic arithmetic expressions, including power and factorial
type CalculatorModule struct{}

// NewCalculatorModule creates a new instance of CalculatorModule
func NewCalculatorModule() *CalculatorModule {
	return &CalculatorModule{}
}

func (c *CalculatorModule) Icon() fyne.Resource {
	icon, err := fyne.LoadResourceFromPath("./icons/calculate.svg")
	if err != nil {
		log.Println("Failed to load icon:", err)
		return nil
	}
	return icon
}

// CanHandle checks if the query contains arithmetic expressions
func (c *CalculatorModule) CanHandle(query string) bool {
	// Regular expression to match numbers, arithmetic operators, and factorial (!), power (^)
	re := regexp.MustCompile(`[\d+\-*/^!()]+`)
	return re.MatchString(query)
}

// Handle evaluates the arithmetic expression and returns the result as a string
func (c *CalculatorModule) Handle(query string) []Result {
	// Preprocess the query to replace '^' with '**' for power, as Go's parser does not recognize '^' for power
	query = strings.ReplaceAll(query, "^", "**")

	// Factorial handling must be done separately as Go's AST doesn't support it natively
	if strings.Contains(query, "!") {
		return []Result{{Title: "Error: Factorial not supported yet", Action: func() {}}}
	}

	expr, err := parser.ParseExpr(query)
	if err != nil {
		return []Result{{Title: "Error: Invalid expression", Action: func() {}}}
	}

	// Evaluate the expression
	result := evalExpr(expr)
	return []Result{{Title: strconv.FormatFloat(result, 'f', -1, 64), Action: func() {}, Icon: c.Icon()}}

}

// evalExpr takes an AST of an expression and recursively evaluates it
func evalExpr(expr ast.Expr) float64 {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.INT || e.Kind == token.FLOAT {
			val, _ := strconv.ParseFloat(e.Value, 64)
			return val
		}
	case *ast.BinaryExpr:
		left := evalExpr(e.X)
		right := evalExpr(e.Y)
		switch e.Op {
		case token.ADD:
			return left + right
		case token.SUB:
			return left - right
		case token.MUL:
			return left * right
		case token.QUO:
			return left / right
		case token.XOR: // Assuming '**' was replaced by '^' and token.XOR is being used for power
			return math.Pow(left, right)
		}
	case *ast.ParenExpr:
		return evalExpr(e.X)
	}
	return 0
}

// TODO: Implement the evaluation of factorial and handle '**' for power explicitly.
