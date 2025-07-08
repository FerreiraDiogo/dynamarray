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

// Reads user input and append it to the slice
func appendToSlice() {
	messages.PrintAppendMessage()
	var element int
	running := true
	for running {

		if !isInputValid(&element) {
			messages.PrintErrorInputMessage()
			continue
		}
		running = false
	}
	dynamarray = append(dynamarray, element)
}

// Reads user input and prepends it to the slice
func prependToSlice() {
	messages.PrintPreppendMessage()
	var element int
	running := true
	for running {

		if !isInputValid(&element) {
			messages.PrintErrorInputMessage()
			continue
		}
		running = false
	}

	var temp = []int{element}
	dynamarray = append(temp, dynamarray...)
}

// Reads user input and validates it. Returns true if input is a valid int,
// false otherwise
func isInputValid(element *int) bool {
	var input int
	_, inputErr := fmt.Scanf("%d", &input)
	fmt.Println(inputErr)
	if inputErr != nil {
		return false
	}
	*element = input
	return true
}
