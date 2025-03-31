package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	list := make([]int, 0, 3)

	fmt.Println("Please enter integers. Enter 'X' to stop.")

	var index int = 1

	for {
		var input string
		fmt.Printf("Please enter %d. integer: ", index)
		fmt.Scan(&input)

		if strings.ToUpper(input) == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		index++

		list = append(list, num)

		sortIntegersAndPrint(list)
	}	

}

func sortIntegersAndPrint(list []int) {
	fmt.Println("You entered the following integers: ", list)
	sort.Ints(list)
	fmt.Println("Sorted integers: ", list)
}