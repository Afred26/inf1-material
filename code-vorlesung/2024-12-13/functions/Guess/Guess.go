package guess

import (
	"fmt"
	"math/rand"
)

func GuessingGame() {
	Zufall := Zufallszahl()
	for i := 0; i < 5; i++ {
		guess := ReadNumber()

		if NumberGood(guess, Zufall) {
			fmt.Println("Richtige geraten! xD")
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

func Zufallszahl() int {
	z := rand.Intn(16) + 1
	return z
}
