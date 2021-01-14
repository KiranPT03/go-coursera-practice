package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Solution(D int) int {
	byteInt := strconv.Itoa(D)
	byteString := strings.Split(byteInt, "")
	var modifiedNumber string
	fiveByteArray := []string{"5"}
	fmt.Println(byteString)
	for index, val := range byteString {
		if val <= "5" {
			fmt.Println(index, val)

			modifiedFiveByteArray := append(fiveByteArray, byteString[index:]...)

			modifiedFiveByteArray = append(byteString[:index], modifiedFiveByteArray...)

			modifiedNumber := strings.Join(modifiedFiveByteArray, "")
			fmt.Println(modifiedNumber)
		}
	}
	fmt.Println(reflect.TypeOf(modifiedNumber))
	value, err := strconv.Atoi(modifiedNumber)
	if err != nil {
		panic(err)
	}
	return int(value)
}

func main() {

	fmt.Println(Solution(0))
}
