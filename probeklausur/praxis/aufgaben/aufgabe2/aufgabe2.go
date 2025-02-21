package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	x1 := -1
	x2 := -1
	for i, el := range list {
		if el == first {
			x1 = i
		}
		if el == last {
			x2 = i
		}

	}

	if x1 >= x2 {
		return []string{}
	}
	result := []string{}
	result = append(result, list[:x1]...)
	result = append(result, list[x2+1:]...)
	return result

}
