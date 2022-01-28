package Files

import (
	"fmt"
)

type Struct struct {
	Name string
	Age  int
}

func DataTypesAndVariables() {

	title := "Go Variables"
	var integer int = 0
	var float float64 = 64
	var boolean bool = true

	person := Struct{
		"Joniyed",
		23,
	}

	fmt.Println(
		title, "\n",
		integer,
		float,
		boolean,
		person,
	)
}
