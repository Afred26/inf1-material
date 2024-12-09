package guess

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

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
}

func NumberGood(n, z int) bool {
	t := false
	if n == z {
		t = true
	}
	return t

}

func Zufallszahl() int {
	z := rand.Intn(5)
	return z
}
