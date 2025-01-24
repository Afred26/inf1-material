package pointer

import (
	"fmt"
)

func Add2(x *int) {
	*x += 2
}

func Example_pointer_function() {
	x := 40
	Add2(&x)

	fmt.Println(x)

	//Output:
	//42
}

type Person struct {
	Name    string
	Vorname string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func Example_pointer_structs() {

	p := Person{"Nym", "Arno"}
	fmt.Println(p)
	p.SetName("Nymous")
	fmt.Println(p)

	//Output:

}

func Example_list() {

	l1 := []int{1, 2, 3, 4, 5}

	fmt.Println(l1)
	Double(l1)
	fmt.Println(l1)
	AppendClone(l1)
	fmt.Println(l1)

	//Output:
	//[1 2 3 4 5]
	//[2 4 6 8 10]
	//[2 4 6 8 10 2 4 6 8 10]
}

// Double erwartet eine Liste und verdoppelt jedes Element
func Double(list []int) {
	for i, el := range list {
		list[i] = el * 2
	}
}

// AppendClone erwartet eine Liste und hängt sie ein zweitesmal an sich selbst
func AppendClone(list []int) {

}

func Example_list_slices() {

	l1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	l2 := l1[3:7]

	fmt.Println(l1)
	fmt.Println(l2)

	l2[2] = 9999

	fmt.Println(l1)
	fmt.Println(l2)

	l2 = append(l2, 15000)

	fmt.Println(l1)
	fmt.Println(l2)

	l2 = append(l2, 15000)

	fmt.Println(l1)
	fmt.Println(l2)

	l2[2] = 4444

	fmt.Println(l1)
	fmt.Println(l2)

	//Output:

}
