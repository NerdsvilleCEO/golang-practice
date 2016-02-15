package main

import "fmt"

func main() {
	str1 := "This"
	str2 := "is"
	str3 := "a test."

	aNumber := 42
	aBool := true

	fmt.Println(str1, str2, str3)
	fmt.Printf("The value of aNumber is: %v\n", aNumber)
	aString := fmt.Sprintf("The value of aBool is: %v\n", aBool)
	fmt.Print(aString)
}
