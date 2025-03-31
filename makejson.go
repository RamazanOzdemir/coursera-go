package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Crreate a map
	person := make(map[string]string)

	name := readProperty("name")
	address := readProperty("address")

	person["name"] = name
	person["address"] = address
	

	jsonPerson := convertMapToJSON(person)


	fmt.Println(jsonPerson)
	
}

func readProperty(property string) string {

	fmt.Printf("Please enter your %s: ", property)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

func convertMapToJSON(person map[string]string) string {

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		return err.Error()
	}
	return string(jsonPerson)
}