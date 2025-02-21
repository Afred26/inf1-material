package hanoi

import "fmt"

func move(s, z string) {
	fmt.Printf("%s -> %s\n", s, z)
}

func Hanoi(s, m, z string, h int) {
	if h == 1 {
		move(s, z)
	} else {
		Hanoi(s, z, m, h-1)
		move(s, z)
		Hanoi(m, s, z, h-1)
	}

}
