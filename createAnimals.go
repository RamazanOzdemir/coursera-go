package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var animalMap = map[string]Animal{}

var commandTypes = map[string]struct{} {
	"newanimal": {},
	"query": {},
}

var animalTypes = map[string]struct{} {
	"cow": {},
	"bird": {},
	"snake": {},
}

var actions = map[string]struct{} {
	"eat": {},
	"move": {},
	"speak": {},
}

var errorMessages = map[string]string{
	"invalidCommand": "The command is not valid.",
	"invalidAnimalType": "The animal type is not valid.",
	"invalidAction": "The action is not valid.",
	"animalAlreadyExists": "The animal name already exists.",
	"invalidInput": "Please provide exactly 3 arguments: <command> <name> <action/type>",
	"invalidAnimal": "The animal does not exist.",
}

type Animal interface {
	Eat()
	Move()
	Speak()
}


type Cow struct {
	Food string
	Locomotion string
	Sound string
}

func (c Cow) Eat() {
	fmt.Printf("The cow eats %s\n", c.Food)
}
func (c Cow) Move() {
	fmt.Printf("The cow moves by %s\n", c.Locomotion)
}
func (c Cow) Speak() {
	fmt.Printf("The cow says %s\n", c.Sound)
}


type Bird struct {
	Food string
	Locomotion string
	Sound string
}

func (b Bird) Eat() {
	fmt.Printf("The bird eats %s\n", b.Food)
}
func (b Bird) Move() {
	fmt.Printf("The bird moves by %s\n", b.Locomotion)
}
func (b Bird) Speak() {
	fmt.Printf("The bird says %s\n", b.Sound)
}


type Snake struct {
	Food string
	Locomotion string
	Sound string
}

func (s Snake) Eat() {
	fmt.Printf("The snake eats %s\n", s.Food)
}
func (s Snake) Move() {
	fmt.Printf("The snake moves by %s\n", s.Locomotion)
}
func (s Snake) Speak() {
	fmt.Printf("The snake says %s\n", s.Sound)
}


type Request interface {
	Validate() (bool, []string)
	Execute() 
}

type NewAnimalRequest struct {
	command string
	name string
	animalType string
}

func (r NewAnimalRequest) Validate() (bool, []string) {
	var errors []string

	// lower the input
	// to make it case insensitive
	r.command = strings.ToLower(r.command)
	r.name = strings.ToLower(r.name)
	r.animalType = strings.ToLower(r.animalType)

	if _, exists := commandTypes[r.command]; !exists {
		errors = append(errors, errorMessages["invalidCommand"])
	}

	if _, exists := animalTypes[r.animalType]; !exists {
		errors = append(errors, errorMessages["invalidAnimalType"])
	}

	if _, exists := animalMap[r.name]; exists {
		errors = append(errors, errorMessages["animalAlreadyExists"])
	}

	return len(errors) == 0, errors
}

func (r NewAnimalRequest) Execute() {
	var animal Animal
	switch r.animalType {
	case "cow":
		animal = Cow{
			Food: "grass",
			Locomotion: "walk",
			Sound: "moo",
		}
	case "bird":
		animal = Bird{
			Food: "worms",
			Locomotion: "fly",
			Sound: "peep",
		}
	case "snake":
		animal = Snake{
			Food: "mice",
			Locomotion: "slither",
			Sound: "hsss",
		}
	}
	animalMap[r.name] = animal
	fmt.Println("Created it!")
}

type QueryRequest struct {
	command string
	name string
	action string
}

func (r QueryRequest) Validate() (bool, []string) {
	var errors []string

	// lower the input
	// to make it case insensitive
	r.command = strings.ToLower(r.command)
	r.name = strings.ToLower(r.name)
	r.action = strings.ToLower(r.action)

	if _, exists := commandTypes[r.command]; !exists {
		errors = append(errors, errorMessages["invalidCommand"])
	}

	if _, exists := animalMap[r.name]; !exists {
		errors = append(errors, errorMessages["invalidAnimal"])
	}

	if _, exists := actions[r.action]; !exists {
		errors = append(errors, errorMessages["invalidAction"])
	}

	return len(errors) == 0, errors
}

func (r QueryRequest) Execute() {
	animal := animalMap[r.name]
	
	switch r.action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
	}
}
	
func GetRequest() (Request, error){
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	requests := strings.Fields(strings.TrimSpace(input))

	if( len(requests) != 3){
		fmt.Println("Please provide exactly 3 arguments: <command> <name> <action/type>")
		return nil, fmt.Errorf("invalid input")
	}

	switch requests[0] {
		case "newanimal":
			return NewAnimalRequest{
				command: requests[0],
				name: requests[1],
				animalType: requests[2],
			}, nil
		case "query":
			return QueryRequest{
				command: requests[0],
				name: requests[1],
				action: requests[2],
			}, nil
		default:
			return nil, fmt.Errorf("invalid command")

	}



}

func main(){
	fmt.Println("Welcome to the Animal Factory!")
	fmt.Println("You can create new animals and query them.")
	fmt.Println("Commands:")
	fmt.Println("newanimal <name> <type>")
	fmt.Println("query <name> <action>")
	fmt.Println("Actions: eat, move, speak")
	fmt.Println("Animal types: cow, bird, snake")
	for{
		fmt.Print("> ")
		request, err := GetRequest()
		if err != nil {
			continue
		}
		valid, errors := request.Validate()
		if !valid {
			for _, error := range errors {
				fmt.Println(error)
			}
			continue
		}
		request.Execute()
	}
}
	