package sorting

func QuickSort(list []int) {
	//Abbruchbedingung
	if len(list) <= 1 {
		return
	}
	//Pivot-Element
	p := list[0]

	//Teillisten
	links := []int{}
	rechts := []int{}

	//Aufteilen
	for _, el := range list[1:] {
		if el < p {
			links = append(links, el)
		} else {
			rechts = append(rechts, el)
		}
	}

	//Rekursion
	QuickSort(links)
	QuickSort(rechts)

	//Zusammenfügen
	result := []int{}
	result = append(result, links...)
	result = append(result, p)
	result = append(result, rechts...)
	//Kopieren
	copy(list, result)
}
