package functions

import (
	"fmt"
	"slices"
)

func Example_sort() {

	l1 := []int{12, 2, 42, 25, 38}
	l2 := []int{12, 2, 42, 25, 38}

	slices.Sort(l1)
	slices.SortFunc(l2, func(a, b int) int {
		return b - a
	})

	fmt.Println(l1)
	fmt.Println(l2)

	//Output:
	//[2 12 25 38 42]
	//[42 38 25 12 2]
}
