package types

import "fmt"

type Length int

func (l Length) Centimeters() int {
	return int(l / 10)
}

func (l Length) Metres() int {
	return l.Centimeters() / 100
}

func (l Length) Kilometres() int {
	return l.Metres() / 1000
}

func FromCentimeters(c int) Length {
	return Length(c * 10)
}

func ExampleLength() {
	//var a Length = 9876543210
	a := FromCentimeters(987654321)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(a.Centimeters())
	fmt.Println(a.Metres())
	fmt.Println(a.Kilometres())

	//Output:
	//9876543210
	//types.Length
	//987654321
	//9876543
	//9876
}
