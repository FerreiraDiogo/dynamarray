package main

import (
	"dynamarray/messages"
	"fmt"
)

const INVALID_INPUT_CHARACTERS = "+_)(*&¨%$#@!\\/|:;>.<,}]^~{[]´`)"

var dynamarray []int

func init() {
	dynamarray = make([]int, 0)
}

func main() {

	running := true
	var input int

	for running {
		messages.PrintWelcomeMessage()

		_, inputErr := fmt.Scanf("%d", &input)
		if inputErr != nil {
			messages.PrintErrorInputMessage()
			continue
		}
		running = selectOption(input)
	}

	messages.PrintGoodbyeMessage()
}

func selectOption(input int) bool {
	switch input {
	case 0:
		return false
	case 1:
		appendToSlice()
		return true
	case 2:
		prependToSlice()
		return true
	case 4:
		messages.PrettyPrint(dynamarray)
		return true
	default:
		messages.PrintErrorInputMessage()
		return true
	}
}

// Appends user input to the slice
func appendToSlice() {
	messages.PrintAppendMessage()

	dynamarray = append(dynamarray, readUserInput())
}

// Prepends the user input to the slice
func prependToSlice() {
	messages.PrintPreppendMessage()

	var temp = []int{readUserInput()}
	dynamarray = append(temp, dynamarray...)
}

// Reads the user input
func readUserInput() int {
	running := true
	var input int
	for running {
		_, inputErr := fmt.Scanf("%d", &input)
		if inputErr != nil {
			messages.PrintErrorInputMessage()
			continue
		}
		running = false
	}
	return input
}
