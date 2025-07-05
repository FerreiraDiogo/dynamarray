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
	case 4:
		messages.PrettyPrint(dynamarray)
		return true
	default:
		messages.PrintErrorInputMessage()
		return true
	}
}

func appendToSlice() {
	messages.PrintAppendMessage()
	var element int
	running := true
	for running {
		_, inputErr := fmt.Scanf("%d", &element)
		if inputErr != nil {
			messages.PrintErrorInputMessage()
			continue
		}
		running = false
	}
	dynamarray = append(dynamarray, element)
}
