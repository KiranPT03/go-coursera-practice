package main

import "fmt"

func main() {
	sli := make([]int, 0, 3)
	sli = append(sli, 100)
	sli = append(sli, 92)
	sli = append(sli, 23)
	sli = append(sli, 123)

	fmt.Println(sli)
	fmt.Println(len(sli))
}
