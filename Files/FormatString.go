package Files

import "fmt"

type Person struct {
	name string
	age  int
	cgpa float64
}

func Format() {

	person := Person{
		"Md. Joniyed Bhuiyan",
		23,
		3.91,
	}

	fmt.Printf("Name: %v, Age: %v, Cgpa: %v", person.name, person.age, person.cgpa)
}
