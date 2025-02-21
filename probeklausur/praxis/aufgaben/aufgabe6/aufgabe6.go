package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	result := []int{}
	for _, el := range l1 {
		if ContainsN(l2, el) {
			result = append(result, el)

		}
	}
	for _, el := range l2 {
		if ContainsN(l1, el) {
			result = append(result, el)
		}
	}
	return result
}
func ContainsN(list []int, n int) bool {
	if len(list) == 0 {
		return false
	}
	for _, el := range list {
		if el == n {
			return false
		}
	}
	return true
}
