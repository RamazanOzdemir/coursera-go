package main

import "fmt"

func GetValue(value *float64 ,valueName string){
	for {
		fmt.Printf("Please enter the %s: ", valueName)
		_, err := fmt.Scan(value)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		break
	}
}

func GetInitialValues() (float64, float64, float64) {
	// get the initial values from the user
	fmt.Println("Please enter the initial values.")
	var acceleration, initialPosition, initialVelocity float64
	GetValue(&acceleration, "acceleration")
	GetValue(&initialPosition, "initial position")
	GetValue(&initialVelocity, "initial velocity")

	return acceleration, initialPosition, initialVelocity
}

func GetTime() float64{
	// get the time from the user
	var time float64
	GetValue(&time, "time")
	return time
}


func GenDisplaceFn(acceleration, initialPosition , initialVelocity float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialPosition 
	}
}

func main() {
	// get the initial values from the user
	acceleration, initialPosition, initialVelocity := GetInitialValues()

	// get the time from the user
	time := GetTime()

	// generate the displacement function
	displaceFn := GenDisplaceFn(acceleration, initialPosition, initialVelocity)

	// calculate the displacement
	displacement := displaceFn(time)

	fmt.Printf("The displacement is: %f\n", displacement)
}
