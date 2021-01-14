package main

import "fmt"

func main() {

	//Map lietrals are defined under curly brackets with comma at the end
	var idMap map[int]string = map[int]string{
		123: "Kiran",
		345: "Tavadare",
	}
	fmt.Println(idMap)

	idMap[2737] = "Django"
	fmt.Println(idMap)

	delete(idMap, 345)
	fmt.Println(idMap)

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
