/*
Question: Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and add the name
and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string
	var person map[string]string = map[string]string{}

	fmt.Println("Please enter the name: ")
	fmt.Scan(&name)

	fmt.Println("Please enter the address: ")
	fmt.Scan(&address)

	person["name"] = name
	person["address"] = address

	objectPerson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(objectPerson)

	jsonStr := string(objectPerson)
	fmt.Println(jsonStr)

}
