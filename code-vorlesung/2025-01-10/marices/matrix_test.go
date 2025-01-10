package marices

import "fmt"

type Matrix [][]float64

func Example_matrix() {
	m := Matrix{
		{1, 3},
		{2, 4},
	}
	fmt.Println(m)

	//Output:
	//[1 3]
	//[2 4]
}

func (m Matrix) String() string {
	var result string

	for _, line := range m {
		result += fmt.Sprintf("%v\n", line)
	}

	return result
}
