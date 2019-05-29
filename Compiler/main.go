package main

import (
	"os"
    "io/ioutil"
    "fmt"
    "github.com/alistairfink/Basic-Go-Compiler/Compiler/DeclarationType"
    "github.com/alistairfink/Basic-Go-Compiler/Compiler/KeyWords"
)

func main() {
	if len(os.Args) < 2 {
        panic(fmt.Sprint("Go File Required"))
	}

	// Take in File Path
	// TODO: Remove all unneccesary prints
	mainFilePath := os.Args[1]
	println("FilePath: ", mainFilePath)
	println()

	// Read File into Memory
	// TODO: Make this open in chunks instead of the entire thing in memory
    file, err := os.Open(mainFilePath)
    if err != nil {
        panic(fmt.Sprint("ERROR: ", err.Error()))
    }

    defer file.Close()
    fileContents, err := ioutil.ReadAll(file)
    if err != nil {
        panic(fmt.Sprint("ERROR: ", err.Error()))
    }

    // Token Parsing
    // TODO: Consider Implementing Regex Matching
    tokens := []string{}
    currToken := ""
    for _, char := range fileContents {
    	if checkSpecialCharacter(string(char)){
    		if len(currToken) > 0 {
    			tokens = append(tokens, currToken)
    			currToken = ""
    		}

    		continue
    	} else if checkValidToken(string(char)) {
    		if len(currToken) > 0 {
				tokens = append(tokens, currToken)
				currToken = ""
			}

			tokens = append(tokens, string(char))
    		continue
    	}

    	currToken += string(char)
    }

    // TODO: Finish other cases
    variables := make(map[string]int)
    var currNode *SyntaxNode
    currNode = nil
    functions := make(map[string]*SyntaxNode)
    for i := 0; i < len(tokens); i++ {
        switch tokens[i] {
        case KeyWords.Function:
            if _, exists := functions[tokens[i+1]]; exists {
                panic(fmt.Sprint("Function Token Already Exists: ", tokens[i+1]))
            }

            newFunction := SyntaxNode{
                name: tokens[i+1],
                declarationType: DeclaractionType.FunctionDeclaration,
                params: make([]interface{}, 0),
                body: []*SyntaxNode{},
                parent: nil,
            }

            currNode = &newFunction
            functions[tokens[i+1]] = &newFunction
            i += 3 
        case KeyWords.Package:
            // TODO: Packages to be implemented
            i++
        case KeyWords.Variable:
            if _, exists := variables[tokens[i+1]]; exists {
                panic(fmt.Sprint("Variable Token Already Exists: ", tokens[i+1]))
            }

            newVar := SyntaxNode{
                declarationType: DeclaractionType.Integer,
                name: tokens[i+1],
                params: make([]interface{}, 0),
                body: []*SyntaxNode{},
                parent: currNode,
            }

            variables[tokens[i+1]] = DeclaractionType.GetValue(tokens[i+2])
            currNode.body = append(currNode.body, &newVar)
            i += 2
        case ")":
            fallthrough
        case "}":
            currNode = currNode.parent
        default:
            if tokens[i + 1] == "(" {
                newCall := SyntaxNode{
                    declarationType: DeclaractionType.Expression,
                    name: tokens[i],
                    params: make([]interface{}, 0),
                    body: []*SyntaxNode{},
                    parent: currNode,
                }

                currNode.body = append(currNode.body, &newCall)
                currNode = &newCall
                i++
            // TODO: Add more operators. Move this somewhere else.
            } else if tokens[i + 1] == "+"{
                operation := SyntaxNode{
                    declarationType: DeclaractionType.Operation,
                    name: tokens[i + 1],
                    params: make([]interface{}, 0),
                    body: []*SyntaxNode{},
                    parent: currNode,
                }

                operation.params = append(operation.params, tokens[i])
                operation.params = append(operation.params, tokens[i+2])
                currNode.body = append(currNode.body, &operation)
                i += 2
            } else if _, exists := variables[tokens[i]]; exists {
                varAssignment := SyntaxNode{
                    declarationType: DeclaractionType.Assignment,
                    name: tokens[i],
                    params: make([]interface{}, 0),
                    body: []*SyntaxNode{},
                    parent: currNode,
                }

                varAssignment.params = append(varAssignment.params, tokens[i + 2])
                currNode.body = append(currNode.body, &varAssignment)
                i += 2
            }
        }
    }

    println("Syntax Tree: ")
    PrintSyntaxTree(functions["main"], 0)

    // TODO: Use AST to generate assembly
    // TODO: Either use shell script or use this to use assembler and then linker to make executables

	PrintSliceString(tokens)
	println()
    println(string(fileContents[:]))
}

func PrintSyntaxTree(root *SyntaxNode, indent int) {
    printable := ""
    for i := 0; i < indent; i++ {
        printable += " "
    }

    printable += "|"

    println(printable + "-----------------------------------")
    println(printable, "Name:", root.name)
    println(printable, "Type:", DeclaractionType.GetString(root.declarationType))
    print(printable, " Params:")
    for _, param := range root.params {
        print(" " + param.(string) + ",")
    }
    println()
    println(printable, "Body:")
    for _, body := range root.body {
        PrintSyntaxTree(body, indent + 1)
    }
}

func PrintSliceString(arr []string) {
	printable := "[ "
	for _, element := range arr {
		printable += element + ", "
	}
	printable = printable[:len(printable)-2] + " ]"
	println(printable)
}