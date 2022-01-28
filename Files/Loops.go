package Files

import "fmt"

func Loops() {

	//while loop
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}

	//for loop
	println()
	for x := 5; x < 15; x++ {
		fmt.Print(x, " ")
	}

	//for each loop
	println()
	names := []string{"Jb", "Sn"}
	for _, name := range names {
		fmt.Print(name, " ")
	}

}
