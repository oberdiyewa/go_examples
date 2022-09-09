package examples

import "fmt"

func ArrayExample() {
	arrInt := [...]int{23, 445, 4564, 34, 56, 6}

	for i := 0; i < len(arrInt); i++ {
		fmt.Println(arrInt[i])
	}

	arrStr := [5]string{"one", "two", "three", "four", "five"}

	for x := 0; x < len(arrStr); x++ {
		fmt.Println(arrStr[x])
	}
}
