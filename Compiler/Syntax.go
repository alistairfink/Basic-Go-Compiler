package main

type SyntaxNode struct {
	name string
	declarationType int
	params []interface{}
	body []*SyntaxNode
	parent *SyntaxNode
}