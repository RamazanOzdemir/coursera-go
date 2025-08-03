package main

import (
	"fmt"
	"sort"
)

// sortArray sorts a sub-array and sends the sorted array back through a channel
func sortArray(arr []int, ch chan []int){
	sort.Ints(arr)
	// Print the sorted sub-array
	fmt.Println("Sorted Sub Array:", arr)
	// Send the sorted array back through the channel
	ch <- arr
}

func main() {
	// Example array to be sorted
	arr := []int{10, 3, 5, 7, 6, 2, 8, 1, 12, 9, 15, 0}

	// Create a channel to receive sorted sub-arrays
	ch := make(chan []int)

	subArr1 := append([]int{}, arr[0:3]...)
	subArr2 := append([]int{}, arr[3:6]...)
	subArr3 := append([]int{}, arr[6:9]...)
	subArr4 := append([]int{}, arr[9:12]...)	
	
	// Start goroutines to sort each sub-array concurrently
	go sortArray(subArr1, ch)
	go sortArray(subArr2, ch)
	go sortArray(subArr3, ch)
	go sortArray(subArr4, ch)

	// Receive sorted sub-arrays from the channel
	sortedSubArr1 := <-ch
	sortedSubArr2 := <-ch
	sortedSubArr3 := <-ch
	sortedSubArr4 := <-ch

	// Merge the sorted sub-arrays
	mergedArr := append(append(sortedSubArr1, sortedSubArr2...), append(sortedSubArr3, sortedSubArr4...)...)
	fmt.Println("Merged Array:", mergedArr)

}