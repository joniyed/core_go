package Files

import (
	"fmt"
	"sort"
	"strings"
)

func BasicLibraries() {
	//string
	str := "I love programming..."
	fmt.Println("Original: ", str)
	fmt.Println(strings.Contains(str, "love"))
	fmt.Println(strings.ReplaceAll(str, "love", "like"))
	fmt.Println(strings.Compare("abc", "aaa"))
	fmt.Println(strings.Count("abc", "aa"))
	fmt.Println(strings.Index(str, "love"))

	//sort
	age := []int{12, 45, 25, 1, 552}
	fmt.Println("Unsorted: ", age)
	sort.Ints(age)
	fmt.Println("Sorted: ", age)
	fmt.Println("After sort search 552 index: ", sort.SearchInts(age, 552))

}
