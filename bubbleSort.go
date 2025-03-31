package main

import "fmt"

func main(){
	list := GetList()
	fmt.Println("Unsorted array: ", list)
	BubbleSort(list)
}


func GetList () []int {

	fmt.Println("Please enter ten integers.")
	list := make([]int, 0, 10)	

	for i:=0; i<10; i++ {
		var input int
		fmt.Printf("Please enter %d. integer: ", i+1)
		_, err :=  fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			i--
			continue
		}

		list = append(list, input)
	}

	return list
}


func BubbleSort(numbers []int) {

	status := map[string]bool{"noChange": true}
	// Loop through the array until no changes are made
	for i:=0; i < len(numbers) -1; i++ {
		CheckNextValue(0, numbers, status)
		
		// If no changes are made end of the tour, break the loop
		if status["noChange"] {
			break
		}

		// Reset the status
		status["noChange"] = true
	}

	fmt.Println("Sorted array: ", numbers)
}

// Swap the values in the array
func Swap( numbers []int, index int){
	temp := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = temp
}

func CheckNextValue(index int, numbers []int, status map[string]bool) {

	// If the index is the last element in the array, break the recursion
	if (index == len(numbers) -1) {
		return
	}

	// If the current value is greater than the next value, swap them
	if numbers[index] > numbers[index+1] {
		Swap(numbers, index)
		status["noChange"] = false
	}

	// Recursively call the function with the next index
	CheckNextValue(index+1, numbers, status)
}



