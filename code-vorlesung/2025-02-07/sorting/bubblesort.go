package sorting

func BubbleSort(list []int) {
	for t := 1; t < len(list); t++ {
		for i := 0; i < len(list)-t; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
	}
}
