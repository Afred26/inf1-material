package rekursion

import "fmt"

func Example_count() {

	CountDown(3)
	CountUp(3)

	//Output:
	//3
	//2
	//1
	//1
	//2
	//3
}

func ExampleSubtract() {
	fmt.Println(Subtract(3))
	fmt.Println(Subtract(4))
	fmt.Println(Subtract(5))

	//Output:
	//2
	//2
	//3
}

func ExamplePotenz() {
	fmt.Println(Potenz(5, 2))
	fmt.Println(Potenz(5, 3))
	fmt.Println(Potenz(5, 4))
	fmt.Println(Potenz(5, 5))
	fmt.Println(Potenz(3, 2))
	fmt.Println(Potenz(3, 3))
	fmt.Println(Potenz(3, 4))
	fmt.Println(Potenz(3, 5))
	fmt.Println(Potenz(3, -2))
	fmt.Println(Potenz(3, -3))

	//Output:
	//25
	//125
	//625
	//3125
	//9
	//27
	//81
	//243
	//0.1111111111111111
	//0.037037037037037035

}

func ExampleFactorial() {

	fmt.Println(Factorial(1))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(5))

	//Output:
	//1
	//2
	//6
	//24
	//120
}
