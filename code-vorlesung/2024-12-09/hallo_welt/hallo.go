package main

import "fmt"

func main() {
	GuessingGame()
}
func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		Zufall := Zufallszahl()
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
	z := ReadNumber()
	return z
}
