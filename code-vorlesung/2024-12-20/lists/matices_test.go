package lists

import "fmt"

func Example_matrix() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("%v \n%v \n%v \n", m1[0], m1[1], m1[2])

	m2 := [][][]int{
		{{1, 2}, {3, 4}},
		{{5, 6}, {7, 8}},
	}
	fmt.Printf("%v \n%v", m2[0], m2[1])

	//Output:
}
