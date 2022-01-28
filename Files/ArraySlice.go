package Files

import "fmt"

func ArraySlices() {
	var ages []int = []int{23, 32, 55}
	var age = ages[0]
	age = 333
	fmt.Println(age)
	var slice = ages[1:2]
	fmt.Println("Array: ", ages, "\nLen: ", len(ages), "\nArray[1:2]: ", slice, "\nLen: ", len(slice))
	ages = append(ages, 77)
	fmt.Println(ages)
}
