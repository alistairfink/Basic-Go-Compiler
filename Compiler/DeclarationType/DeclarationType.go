package DeclaractionType

import(
	"fmt"
)

const (
	FunctionDeclaration	= 0
	Expression			= 1
	Assignment			= 2
	Operation			= 3
	Integer				= 4
)

func GetValue(name string) int {
	switch(name) {
	case "FunctionDeclaration":
		return FunctionDeclaration
	case "Expression":
		return Expression
	case "Assignment":
		return Assignment
	case "Operation":
		return Operation
	case "Integer":
		fallthrough
	case "int":
		return Integer
	default:
		panic(fmt.Sprint("Invalid Declartion Type: ", name))
	}
}