package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkStartWithI(input string) bool {
	return strings.HasPrefix(input, "i")
}

func checkEndWithN(input string) bool {
	return strings.HasSuffix(input, "n")
}

func checkContainA(input string) bool {
	return strings.Contains(input, "a")
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if checkStartWithI(input) && checkEndWithN(input) && checkContainA(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}


}