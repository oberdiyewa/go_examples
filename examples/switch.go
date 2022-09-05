package examples

import "fmt"

func SwitchExample() {
	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case number < 11:
		fmt.Println("Number < 11")
		fallthrough
	default:
		fmt.Println("Default message")
	}
}
