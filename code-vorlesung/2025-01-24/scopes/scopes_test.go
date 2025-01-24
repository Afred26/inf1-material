package scopes

import "fmt"

func Example_scope() {
	x := Foo(3)

	fmt.Println(x)

	//Output:
	//70

}

func Foo(t int) int {
	x := 42
	Bar(&x)
	y := 23
	z := x + y + t
	return z
}

func Bar(x *int) {
	*x += 2
}

func Example_scope_2() {

	x := 10
	{
		y := 20
		{
			z := 30
			fmt.Println(z)
			fmt.Println(y)
			fmt.Println(x)
		}
		fmt.Println(y)
	}
	fmt.Println(x)

	//Output:
	//30
	//20
	//10
	//20
	//10
}

func Example_scope_3() {

	i := 0

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	//Output:
	//0
	//1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//0

}

func Example_scope_4() {

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	//Output:
	//0
	//1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10
}
