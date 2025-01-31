package functions

import "fmt"

// ForEach erwartet eine Liste aus Zahlen und eine Funktion.
// Wendet f auf jedes Element an.
func ForEach(list []int, f func(int)) {
	for _, el := range list {
		f(el)
	}
}

func PrintNum(n int) {

	fmt.Println(n)
}

func PrintTwice(n int) {
	fmt.Println(n)
	fmt.Println(n)
}

func ExampleForEach() {

	l1 := []int{5, 2, 4, 7, 8, 42}

	ForEach(l1, PrintNum)
	ForEach(l1, PrintTwice)

	//Output:

}

func Example_count() {

	l1 := []int{5, 2, 4, 7, 8, 42}

	sum1 := 0
	sum2 := 0
	add := func(n int) {
		sum1 += n
	}

	sum_evan := func(n int) {
		if n%2 == 0 {
			sum2 += n
		}
	}

	ForEach(l1, add)
	ForEach(l1, sum_evan)

	fmt.Println(sum1)
	fmt.Println(sum2)

	//Output:
	//68
	//56
}

func Example_filter_evan() {

	l1 := []int{5, 2, 4, 7, 8, 31, 42}

	l2 := []int{}

	filter_evan := func(n int) {
		if n%2 == 1 {
			l2 = append(l2, n)
		}
	}

	ForEach(l1, filter_evan)

	fmt.Println(l2)

	//Output:
	//[5 7 31]
}
