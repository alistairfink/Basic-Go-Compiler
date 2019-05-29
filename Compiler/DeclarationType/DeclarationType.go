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
		panic(fmt.Sprint("Invalid Declartion Type:", name))
	}
}

func GetString(number int) string {
	switch(number) {
	case 0:
		return "FunctionDeclaration"
	case 1:
		return "Expression"
	case 2:
		return "Assignment"
	case 3:
		return "Operation"
	case 4:
		fallthrough
	case 5:
		return "Integer"
	default:
		panic(fmt.Sprint("Invalid Declartion Type:", number))
	}
}