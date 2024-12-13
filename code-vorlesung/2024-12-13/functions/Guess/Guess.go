package guess

import (
	"fmt"
	"math"
	"math/rand"
)

func GuessingGame() {

	var s, länge int
	fmt.Println("Wie lange möchten Sie spielen?")
	fmt.Scanln(&länge)
	fmt.Println("Wie schwierig soll es sein?")
	fmt.Scanln(&s)
	Zufall := Zufallszahl(länge, s)
	for i := 0; i < länge; i++ {
		guess := ReadNumber()

		if NumberGood(guess, Zufall) {
			return
		}
	}
	fmt.Println("Zuviele falsche Zahlen! :(")
	fmt.Println("Die richtige Zahl war", Zufall)
}

func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

func NumberGood(guess, Zufall int) bool {

	if guess == Zufall {
		fmt.Println("Richtige geraten! xD")
		return true
	} else {
		if guess > Zufall {
			fmt.Println("Zahl zu groß")
		}
		if guess < Zufall {
			fmt.Println("Zahl zu klein")
		}
	}
	return false

}

func Zufallszahl(länge, s int) int {
	q := float64(länge + s)
	p := int(math.Pow(2, q))
	z := rand.Intn(p) + 1
	return z
}
