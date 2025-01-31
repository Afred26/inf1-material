package functions

import "fmt"

func PrintFoo() {
	fmt.Println("Foo")
}

func PrintBar() {
	fmt.Println("Bar")
}

// CallFunction erwartet einen Funktion als Parameter und ruft sie auf.
// Die Funktion, die erwartet wird, hat selbst keine Parameter und liefert nichts zurück
func CallFunction(f func()) {
	f()
}

func ExamplePrintFoo() {

	PrintBaz := func() {
		fmt.Println("Baz")
	}
	CallFunction(PrintBar)
	CallFunction(PrintFoo)
	CallFunction(PrintBaz)

	f1 := CreatePrinter("ABC")
	CallFunction(f1)
	CallFunction(CreatePrinter("CBA"))

	//Output:
	//Bar
	//Foo
	//Baz
	//ABC
	//CBA
}

func CreatePrinter(s string) func() {
	return func() {
		fmt.Println(s)
	}
}
