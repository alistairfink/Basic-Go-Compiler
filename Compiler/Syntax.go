package main

const (
	DeclaractionType = map[string]int{
		"FunctionDeclaration": 0,
		"VariableDeclaration": 1,
		"Expression": 2,		
	}
)

type SyntaxNode struct {
	declarationType int
	params []interface{}
	body []SyntaxNode
}