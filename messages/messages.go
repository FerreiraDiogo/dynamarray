package messages

import "fmt"

func PrintWelcomeMessage() {
	fmt.Println("==========Dynamarray V1.0.0==========")
	fmt.Println("==========     Options     ==========")
	fmt.Print("0 - Quit Program\n1 - Append Integer to Slice\n2 - Prepend Integer to Slice\n3 - Insert Integer into especified position\n" +
		"4 - List all integers\n5 - Find integer\n6 - Remove integer\n7 - Reorder\n")
}
