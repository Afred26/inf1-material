package structs

import "fmt"

func Example_zugriff() {
	p1 := NewPerson("Beispiel", "Berta", 23)

	fmt.Println(p1)
	p1.Alter = 77
	fmt.Println(p1)
	p1.Name = ("Musterman")
	fmt.Println(p1)
	p1.Vorname = ("Max")
	fmt.Println(p1)
	p1.Phone = PhoneNummber{CountryCode: "+49", AreaCode: "176", Number: "1234567"}
	fmt.Println(p1)

	//Output:
	//{Beispiel Berta 23 {  }}
	//{Beispiel Berta 77 {  }}
	//{Musterman Berta 77 {  }}
	//{Musterman Max 77 {  }}
	//{Musterman Max 77 {+49 176 1234567}}
}
