package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	result := []int{}
	if len(l1) == 0 {
		return append(result, l2...)
	}
	if len(l2) == 0 {
		return append(result, l1...)
	}
	if l1[0] > l2[0] {
		result = append(result, l1[0])
	} else {
		result = append(result, l2[0])
	}
	return append(result, MaxElements(l1[1:], l2[1:])...)
}
