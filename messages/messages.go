package messages

import "fmt"

func PrintWelcomeMessage() {
	fmt.Println("==========Dynamarray V1.0.0==========")
	fmt.Println("==========     Options     ==========")
	fmt.Print("0 - Quit Program\n1 - Append Integer to Slice\n2 - Prepend Integer to Slice\n3 - Insert Integer into especified position\n" +
		"4 - List all integers\n5 - Find integer\n6 - Remove integer\n7 - Reorder\n")
}

func PrintGoodbyeMessage() {
	fmt.Println("Thanks for using Dynamarray!")
}

func PrintErrorInputMessage() {
	fmt.Println("Please, type a valid option")
}

func PrettyPrint(list []int) {
	fmt.Print("The elements in the slice are: [ ")
	for _, v := range list {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("]")
}

func PrintAppendMessage() {
	fmt.Println("Insert the value to be appended to the slice")
}

func PrintPreppendMessage() {
	fmt.Println("Insert the value to be prepended to the slice")
}

func PrintInsertValueMessage() {
	fmt.Println("Insert the value to be inserted in the slice.")
}

func PrintInsertPositionValueMesage() {
	fmt.Println("Now type the index to which it should be inserted. Please note that the element in the selected index will be overwriten.")
}
func PrintInvalidPositionValueMessage() {
	fmt.Println("You can't insert beyond the slice's length. Your input will be appended")
}
func PrintRemoveValueMessage() {
	fmt.Println("Remove by index or value?\n1 - Index\n2 - Value")
}

func PrintRemoveValueByIndex() {
	fmt.Println("Input a single index to remove the element at that position,  or a couple to remove all itens in the range, last one included.")
}

func PrintRemoveValueIfExists() {
	fmt.Println("Input the value you want to remove.You will be warned if it doesn't exists in the slice.")
}

func PrintInexistentValueMessage() {
	fmt.Println("Inexistent value")
}
