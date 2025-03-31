package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Type string
	Food string
	Locomotion string
	Sound string
}

var animals map[string]Animal = map[string]Animal{
	"cow": Animal{Type: "cow", Food: "grass", Locomotion: "walk", Sound: "moo"},
	"bird": Animal{Type: "bird", Food: "worms", Locomotion: "fly", Sound: "peep"},
	"snake": Animal{Type: "snake", Food: "mice", Locomotion: "slither", Sound: "hsss"},
}

var animalTypes map[string]string = map[string]string{
	"cow": "cow",
	"bird": "bird",
	"snake": "snake",}

var actions map[string]string = map[string]string{
	"eat": "eat",
	"move": "move",
	"speak": "speak",}

func (animal Animal) Eat() {
	fmt.Printf("%s eats %s.\n", animal.Type, animal.Food)
}

func (animal Animal) Move() {
	fmt.Printf("%s moves by %s.\n", animal.Type, animal.Locomotion)
}

func (animal Animal) Speak() {
	fmt.Printf("%s says %s.\n", animal.Type, animal.Sound)
}

func GetRequest() (string, string) {
	// get the animal type and action from the user
	var animalType, action string
	fmt.Print(">")
	fmt.Scanln(&animalType, &action)
	return animalType, action
}

func ValidateRequest(animalType, action string) (bool , []string){
	var valid bool = true

	// lower the input
	// to make it case insensitive
	animalType = strings.ToLower(animalType)
	action = strings.ToLower(action)

	// check if the animal type and action are valid
	// if not, set valid to false and add the error message to the errors slice
	var errors []string
	if _, ok := animalTypes[animalType]; !ok {
		valid = false
		errors = append(errors, fmt.Sprintf("The animal type %s is not valid.", animalType))
	}

	if _, ok := actions[action]; !ok {
		valid = false
		errors = append(errors, fmt.Sprintf("The action %s is not valid.", action))
	}

	return valid, errors

}

func main(){
	fmt.Println("Please enter the name of the animal and the action you want to perform (eat, move, speak).")
	
	for{
		animalType, action := GetRequest()
		valid, errors := ValidateRequest(animalType, action)
		if !valid {
			for _, error := range errors {
				fmt.Println(error)
			}
			continue
		}
		animal := animals[animalType]

		switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
		}

	}
	
	
}