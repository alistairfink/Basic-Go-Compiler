package main

import (
	"os"
    "io/ioutil"
)

func main() {
	if len(os.Args) < 2 {
		println("Go File Required")
		return
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
		println("ERROR: ", err.Error())
    }

    defer file.Close()
    fileContents, err := ioutil.ReadAll(file)
    if err != nil {
		println("ERROR: ", err.Error())
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

	PrintSliceString(tokens)
	println()
    println(string(fileContents[:]))
}

func PrintSliceString(arr []string) {
	printable := "[ "
	for _, element := range arr {
		printable += element + ", "
	}
	printable = printable[:len(printable)-2] + " ]"
	println(printable)
}