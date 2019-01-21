package main

import (
	. "./FunctionalUtils"
	"fmt"
)

func main() {

	list := []int{81, 32, 53, 14, 75, 96, 37, 58, 29, 10}

	//ascending := func(a, b int) bool { return a < b }
	descending := func(a, b int) bool { return a > b }
	
	sorted := SortInt(descending, list)

	fmt.Println(list)
	fmt.Println(sorted)

}