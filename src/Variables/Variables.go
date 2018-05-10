package main

import "fmt"
import "strconv"

func main() {
	// Declaracion
	var age = 20
	var age2 string

	age2 = "2"

	// Declaracion y asignacion
	agestring := strconv.Itoa(age)

	fmt.Println("Que tal tengo " + agestring + " y " + age2)
}
