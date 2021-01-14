/*
Question: Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var enteredString string
	fmt.Println("Please enter the string: ")
	fmt.Scanln(&enteredString)
	fmt.Println(enteredString)
	enteredString = strings.ToLower(enteredString)
	enteredString = strings.TrimSpace(enteredString)
	if strings.Contains(enteredString, "i") && strings.Contains(enteredString, "a") && strings.Contains(enteredString, "n") {
		byteString := strings.Split(enteredString, "")
		fmt.Println(byteString)
		if byteString[0] == "i" && byteString[len(byteString)-1] == "n" {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}

}
