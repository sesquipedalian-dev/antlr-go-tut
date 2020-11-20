package main

import (
	"fmt"
	"strconv"

	"github.com/sesquipedalian-dev/antlr-go-tut/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	StateNode = 1
	StateEdge = 2
)

type helloListener struct {
	*parser.BasehelloListener

	currentNode int
	state       int
	result      BfsInput
}

// EnterExpression is called when entering the expression production.
// EnterExpression(c *ExpressionContext)
// ExitExpression(c *ExpressionContext)
// ExitEdges(c *EdgesContext)

func NewHelloListener() *helloListener {
	out := helloListener{
		result: BfsInput{
			Edges: make(map[int][]int),
		},
	}
	return &out
}

func (l *helloListener) EnterExpression(c *parser.ExpressionContext) {
	l.state = StateNode
}

func (l *helloListener) EnterEdgeExpr(c *parser.EdgeExprContext) {
	l.state = StateEdge
}

func (l *helloListener) ExitId(c *parser.IdContext) {
	nodeIDStr := c.GetText()
	nodeID, err := strconv.Atoi(nodeIDStr)
	if err != nil {
		panic(err.Error())
	}

	switch l.state {
	case StateNode:
		l.result.Nodes = append(l.result.Nodes, nodeID)
		l.currentNode = nodeID
	case StateEdge:
		currentEdges, ok := l.result.Edges[l.currentNode]
		if !ok {
			currentEdges = []int{}
		}

		l.result.Edges[l.currentNode] = append(currentEdges, nodeID)

		// no-op
	}
}

// func (l *helloListener) ExitExpression(c *parser.ExpressionContext) {
// 	nodeIDStr := c.GetNodeID().GetText()
// 	nodeID, err := strconv.Atoi(nodeIDStr)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	l.result.Nodes = append(l.result.Nodes, nodeID)

// 	c.GetEdges
// }

func main() {
	// Setup the input
	is, err := antlr.NewFileStream("hello_test.txt")
	if err != nil {
		panic(fmt.Sprintf("Can't open file %s", err))
	}

	// Create the Lexer
	lexer := parser.NewhelloLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewhelloParser(stream)

	// Finally parse the expression
	listener := NewHelloListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	fmt.Printf("output: %s\n", listener.result)
}
