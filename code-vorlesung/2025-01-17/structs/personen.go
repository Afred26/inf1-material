package structs

type Personen struct {
	Name    string
	Vorname string
	Alter   int
	Phone   PhoneNummber
}

type PhoneNummber struct {
	CountryCode string
	AreaCode    string
	Number      string
}

func NewPerson(name string, Prename string, age int) Personen {
	return Personen{Name: name, Alter: age, Vorname: Prename, Phone: PhoneNummber{}}
}
