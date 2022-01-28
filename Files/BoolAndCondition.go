package Files

import (
	"fmt"
	"strings"
)

func BoolAndCondition() {

	pass := true
	mark := 10

	for i := 1; mark*i < 100; i++ {

		mark *= i
		if mark < 33 {
			pass = false
		} else {
			pass = true
		}

		if pass {
			fmt.Println("Passed", mark)
		} else {
			fmt.Println("Failed", mark)
		}
	}

	names := []string{"JB", "SN", "OTHER"}

	for _, name := range names {
		if strings.EqualFold(name, "OTHER") {
			continue
		}
		fmt.Println(name)
	}

}
