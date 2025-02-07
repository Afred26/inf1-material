package sorting

import "fmt"

func ExampleBubbleSort() {
	l1 := []int{12, 2, 42, 25, 38, 56, 78, 1, 99, 23}
	BubbleSort(l1)
	fmt.Println(l1)
	// Output:
	// [1 2 12 23 25 38 42 56 78 99]
}
