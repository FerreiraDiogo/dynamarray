package main

import (
	"bufio"
	"dynamarray/enum"
	"dynamarray/messages"
	"dynamarray/sorter"
	"dynamarray/statistics"
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
var sortState enum.SortState

func init() {
	dynamarray = make([]int, 0, 100)
	bufferedReader = bufio.NewReader(os.Stdin)
	sortState = enum.Unknown
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

// Prints all the program's options
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
		messages.PrintSorting(sortState)
		return true
	case 5:
		findElement()
		return true
	case 6:
		removeElement()
		return true
	case 7:
		reorder()
		return true
	case 8:
		messages.PrintSorting(sortState)
		return true
	case 9:
		showStatistics()
		return true
	default:
		messages.PrintErrorInputMessage()
		return true
	}
}

// Shows statistics functions that are run on the slices elements
func showStatistics() {
	if len(dynamarray) != 0 {
		if sortState == enum.Unknown {
			messages.PrintSorting(sortState)
			reorder()
		}
		meanValue, meanErr := statistics.Mean(dynamarray)
		medianValue, medianErr := statistics.Median(dynamarray)
		modeKey, modeValue, modeErr := statistics.Mode(dynamarray)

		if meanErr != nil || medianErr != nil || modeErr != nil {
			messages.PrintErrorMessages(meanErr, medianErr, modeErr)
		} else {
			messages.PrintStatisticData(meanValue, medianValue, modeKey, modeValue)
		}
	} else {
		messages.PringSliceIsEmpty()
	}
}

// finds and returns an element and its indexs. If its not found, prints a warning message
func findElement() {
	if len(dynamarray) != 0 {
		if sortState == enum.Unknown {
			messages.PrintSorting(sortState)
			reorder()
		}
		messages.PrintSearchValue()
		input := readUserInput()
		index, found := slices.BinarySearch(dynamarray, input)
		if !found {
			messages.PrintNoValueFound()
		} else {
			messages.PrintValuesFound(input, index)
		}

	} else {
		messages.PringSliceIsEmpty()
	}
}

// sorts the array on ascending or descending order, given by user input
func reorder() {
	messages.PrintSortMessage()
	option := readUserInput()
	switch option {
	case 1:
		sorter.QuickSort(dynamarray, 0, len(dynamarray)-1, sorter.PartitionSortAscending)
		messages.PrettyPrint(dynamarray)
		setSortState(enum.SortAscending)
	case 2:
		sorter.QuickSort(dynamarray, 0, len(dynamarray)-1, sorter.PartitionSortDescending)
		messages.PrettyPrint(dynamarray)
		setSortState(enum.SortDescending)
	}
}

// sets the sortState flag
func setSortState(state enum.SortState) {
	sortState = state
	messages.PrintSorting(sortState)
}

// gives options to user to select which way to remove elements from the slice
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

// wrapper function for removal of values from the slice
func removeFromSlice(bufferedReader *bufio.Reader, strategy func(value []int) []int) []int {
	input, err := readMultipleInput(bufferedReader)
	if err != nil {
		fmt.Println(err)
		return dynamarray
	} else {
		return strategy(input)
	}
}

// removes a value from the slice on the given index if it exists, prints a warning message otherwise
func removeByIndexes(indexesSlice []int) []int {
	if len(indexesSlice) == 1 {
		return slices.Delete(dynamarray, indexesSlice[0], indexesSlice[0]+1)
	}
	return slices.Delete(dynamarray, indexesSlice[0], indexesSlice[1]+1)
}

// removes a value from the slice if it exists, prints a warning message otherwise
func removeByValue(value []int) []int {
	for i, v := range dynamarray {
		if v == value[0] {
			return slices.Delete(dynamarray, i, i+1)
		}
	}
	messages.PrintInexistentValueMessage()
	return dynamarray
}

// reads multiple values in a single input
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
		conv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid input: %s is not an integer. Skipping.\n", v)
			continue
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
