package main

import (
	"bufio"
	"dynamarray/messages"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INVALID_INPUT_CHARACTERS = "+_)(*&¨%$#@!\\/|:;>.<,}]^~{[]´`)"

var dynamarray []int
var bufferedReader *bufio.Reader

func init() {
	dynamarray = make([]int, 0, 100)
	bufferedReader = bufio.NewReader(os.Stdin)
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
	case 3:
		insertInPosition()
		return true
	case 4:
		messages.PrettyPrint(dynamarray)
		return true
	case 6:
		removeElement()
		return true
	default:
		messages.PrintErrorInputMessage()
		return true
	}
}

func removeElement() {
	messages.PrintRemoveValueMessage()
	option := readUserInput()
	switch option {
	case 1:
		messages.PrintRemoveValueByIndex()
		dynamarray = removeFromSlice(bufferedReader, removeByIndexes)
	case 2:
		messages.PrintRemoveValueIfExists()
		dynamarray = removeFromSlice(bufferedReader, removeByValue)
	}

}

func removeFromSlice(bufferedReader *bufio.Reader, strategy func(value []int) []int) []int {
	input, err := readMultipleInput(bufferedReader)
	if err != nil {
		fmt.Println(err)
		return dynamarray
	} else {
		return strategy(input)
	}
}

func removeByIndexes(indexesSlice []int) []int {
	if len(indexesSlice) == 1 {
		return slices.Delete(dynamarray, indexesSlice[0], indexesSlice[0]+1)
	}
	return slices.Delete(dynamarray, indexesSlice[0], indexesSlice[1]+1)
}

func removeByValue(value []int) []int {
	for i, v := range dynamarray {
		if v == value[0] {
			return slices.Delete(dynamarray, i, i+1)
		}
	}
	messages.PrintInexistentValueMessage()
	return dynamarray
}

func readMultipleInput(reader *bufio.Reader) ([]int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	strNumbers := strings.Fields(input)

	if len(strNumbers) > 2 {
		return nil, errors.New("type at most 2 values")
	}

	convertedInput := make([]int, 0, len(strNumbers))
	for _, v := range strNumbers {
		conv, _ := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid input: %s is not an integer. Skipping.\n", v)
			continue // Skip invalid entries
		}
		convertedInput = append(convertedInput, conv)
	}
	return convertedInput, nil
}

// inserts user input in user specified position
func insertInPosition() {
	messages.PrintInsertValueMessage()
	valueInput := readUserInput()
	messages.PrintInsertPositionValueMesage()
	positionInput := readUserInput()
	insert(valueInput, positionInput)
}

// Inserts a value into specified position, overwriting it's content. If
// position is beyond slice's length, appends the value to the slice.
func insert(valueInput, positionInput int) {

	if positionInput >= len(dynamarray) {
		messages.PrintInvalidPositionValueMessage()
		dynamarray = append(dynamarray, valueInput)
	} else {
		dynamarray[positionInput] = valueInput

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
